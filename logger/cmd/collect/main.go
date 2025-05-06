package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	slog "log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/IBM/sarama"
	"github.com/Kifiya-Financial-Technology/Logger/handler"
	"github.com/Kifiya-Financial-Technology/Logger/internal/application"
	"github.com/Kifiya-Financial-Technology/Logger/internal/application/commands"
	"github.com/Kifiya-Financial-Technology/Logger/internal/domain"
	"github.com/Kifiya-Financial-Technology/Logger/internal/postgres"
	"github.com/Kifiya-Financial-Technology/Logger/loggerpb"
	"github.com/Kifiya-Financial-Technology/Logger/lswagger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pg "gorm.io/driver/postgres"
    _ "github.com/lib/pq" // 👈 required for sql.Open("postgres", ...)

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gorm.io/gorm"
)

const (
	topic         = "logs"
	consumerGroup = "logger-service-group" // Use consumer groups for scalability
	numPartitions = 3                      // Match your Kafka topic partitions
	batchSize     = 100                    // Process logs in batches
	flushInterval = 1 * time.Second        // Max delay before flushing a batch
	maxRetries    = 3                      // Retry failed messages
	kafkaBrokers  = "kafka1:9091"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		slog.Println("⚠️ No .env file found or failed to load it")
	} else {
		slog.Println("✅ .env file loaded")
	}

	port := os.Getenv("PORT")
	grpcPort := os.Getenv("GRPC_PORT")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DATABASE")

	db, err := init_database(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		slog.Fatalf("❌ failed to connect to DB: %v", err)
	}

	repo := postgres.NewLogRepository(db)
	app := application.New(repo)

	grpcAddress := fmt.Sprintf("0.0.0.0:%s", grpcPort)

	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		slog.Fatalf("❌ failed to listen on port %v: %v", grpcPort, err)
	}
	grpcServer := grpc.NewServer()
	handler.NewServer(app, grpcServer)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	httpMux := http.NewServeMux()
	httpMux.Handle("/", mux)
	httpMux.Handle("/swagger/", lswagger.SwaggerHandler())
	err = loggerpb.RegisterLoggerServiceHandlerFromEndpoint(
		ctx,
		mux,
		grpcAddress,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		slog.Fatalf("❌ failed to register gRPC handlers: %v", err)
	}

	fmt.Printf("🌐 HTTP server (REST + Swagger) running on :%s\n", port)
	go func() {
		if err := http.ListenAndServe(":"+port, httpMux); err != nil {
			slog.Fatalf("❌ HTTP server failed: %v", err)
		}
	}()

	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Fetch.Min = 1 << 20     // 1MB
	config.Consumer.Fetch.Default = 5 << 20 // 5MB
	config.ChannelBufferSize = 1024         // Buffer size

	consumer, err := sarama.NewConsumer([]string{kafkaBrokers}, config)
	if err != nil {
		slog.Fatalf("❌ failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		slog.Fatalf("❌ failed to get list of partitions: %v", err)
	}

	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			slog.Printf("❌ Failed to consume partition %d: %v", partition, err)
			continue
		}

		go func(pc sarama.PartitionConsumer) {
			defer pc.Close()

			for message := range pc.Messages() {
				var log domain.Log
				// Attempt to unmarshal the Kafka message into a log struct
				if err := json.Unmarshal(message.Value, &log); err != nil {
					// If unmarshalling fails, print error and continue to the next message
					fmt.Printf("⚡ Invalid Kafka message: %s\n", string(message.Value))
					continue
				}

				// Process the log by creating a new log entry in your system
				_, err := app.CreateLog(
					ctx,
					commands.CreateLogCommand{
						FromApplication: log.FromApplication,
						Level:           log.Level,
						DateTime:        log.DateTime,
						Message:         log.Message,
					})

				// Check if there was an error during log creation
				if err != nil {
					slog.Printf("⚠️ Failed to process log: %v\n", err)
				} else {
					// Successfully processed log
					slog.Printf("✅ Log processed: %s\n", log.Message) // print the actual message, not the raw value
				}

				// You might want to commit the offset after processing each message if you use manual offset management
				// Example: pc.MarkOffset(message, "") // Marks the message offset (for manual offset management)
			}
		}(pc)
	}
	slog.Fatal(grpcServer.Serve(listener))

}

func init_database(dbHost, dbPort, dbUser, dbPass, dbName string) (*gorm.DB, error) {
	tempDSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", dbHost, dbPort, dbUser, dbPass)
	sqlDB, err := sql.Open("postgres", tempDSN)
	if err != nil {
		log.Fatalf("❌ Unable to connect to Postgres for DB check: %v", err)
		return nil, err
	}
	defer sqlDB.Close()

	var exists bool
	checkQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)
	if err := sqlDB.QueryRow(checkQuery).Scan(&exists); err != nil {
		log.Fatalf("❌ Failed to check if DB exists: %v", err)
		return nil, err
	}

	if !exists {
		_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			log.Fatalf("❌ Failed to create database %s: %v", dbName, err)
			return nil, err
		}
		log.Printf("✅ Created database %s\n", dbName)
	} else {
		log.Printf("✅ Database %s already exists\n", dbName)
	}

	// Connect using GORM to the actual app database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
		return nil, err
	}
	log.Println("✅ Database connected")

	// Enable uuid-ossp extension
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Fatalf("❌ Failed to enable uuid-ossp extension: %v", err)
		return nil, err
	}

	// Auto-migrate the models
	if err := db.AutoMigrate(&domain.Log{}); err != nil {
		log.Fatalf("❌ Failed to auto-migrate: %v", err)
		return nil, err
	}
	log.Println("✅ Auto-migration complete")

	return db, nil
}
