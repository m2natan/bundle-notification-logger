package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/IBM/sarama"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application/commands"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application/queries"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
)

type Application struct {
	application.App
	logger *slog.Logger
}

type LogLevel string

var brokers = []string{"kafka1:9091"}

func createLogKafkaProducer(Data Logger) error {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true // Needed for SyncProducer

	// Create SyncProducer (better for waiting for ack)
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("❌ Failed to create Kafka producer: %v", err)
		return err
	}
	defer producer.Close()

	topic := "logs"

	data, err := json.Marshal(Data)
	if err != nil {
		log.Fatalf("❌ Failed to marshal notification: %v", err)
		return err

	}

	// Create ProducerMessage
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   nil,
		Value: sarama.ByteEncoder(data),
	}

	// Send Message
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("❌ Failed to send message: %v", err)
		return err

	}
	log.Printf("✅ Message sent: partition=%d, offset=%d", partition, offset)

	log.Println("✅ All Log data are sent to Kafka!")
	// Optional: sleep to allow Kafka to flush
	time.Sleep(1 * time.Second)
	return nil
}

const (
	LevelInfo    LogLevel = "INFO"
	LevelWarning LogLevel = "WARNING"
	LevelError   LogLevel = "ERROR"
	LevelDebug   LogLevel = "DEBUG"
	LevelFatal   LogLevel = "FATAL"
	LevelUnknown LogLevel = "UNKNOWN"
)

type Logger struct {
	Level           LogLevel `json:"level"`
	FromApplication string   `json:"source"`
	DateTime        string   `json:"date_time"`
	Message         string   `json:"message"`
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(application application.App, logger *slog.Logger) Application {
	return Application{
		App:    application,
		logger: logger,
	}
}

func SerializeLogMessage(msg Logger) error {
	return createLogKafkaProducer(msg)
}

func (a Application) FindAll(ctx context.Context) (notification []domain.Notification, err error) {
	a.logger.Info("--> Notifications.FindAll")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.FindAll", "error", err)
			SerializeLogMessage(Logger{
				Level:           LevelError,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.FindAll error: %s", err),
			})
		} else {
			a.logger.Info("<-- Notifications.FindAll")
			SerializeLogMessage(Logger{
				Level:           LevelInfo,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.FindAll"),
			})
		}
	}()

	return a.App.FindAll(ctx)
}

func (a Application) FindById(ctx context.Context, query queries.FindById) (notification *domain.Notification, err error) {
	a.logger.Info("--> Notifications.FindById")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.FindById", "error", err)
			SerializeLogMessage(Logger{
				Level:           LevelError,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.FindById error: %s", err),
			})
		} else {
			a.logger.Info("<-- Notifications.FindById")
			SerializeLogMessage(Logger{
				Level:           LevelInfo,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.FindById"),
			})
		}
	}()

	return a.App.FindById(ctx, query)
}

func (a Application) FindByStatus(ctx context.Context, query queries.FindByStatus) (notification []domain.Notification, err error) {
	a.logger.Info("--> Notifications.FindByStatus")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.FindByStatus", "error", err)
			SerializeLogMessage(Logger{
				Level:           LevelError,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.FindByStatus error: %s", err),
			})
		} else {
			a.logger.Info("<-- Notifications.FindByStatus")
			SerializeLogMessage(Logger{
				Level:           LevelInfo,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.FindByStatus"),
			})
		}
	}()

	return a.App.FindByStatus(ctx, query)
}

func (a Application) FindByType(ctx context.Context, query queries.FindByType) (notification []domain.Notification, err error) {
	a.logger.Info("--> Notifications.FindByType")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.FindByType", "error", err)
			SerializeLogMessage(Logger{
				Level:           LevelError,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.FindByType error: %s", err),
			})
		} else {
			a.logger.Info("<-- Notifications.FindByType")
			SerializeLogMessage(Logger{
				Level:           LevelInfo,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.FindByType"),
			})
		}
	}()

	return a.App.FindByType(ctx, query)
}

func (a Application) CreateNotification(ctx context.Context, cmd commands.CreateNotificationCommand) (notification *domain.Notification, err error) {
	a.logger.Info("--> Notifications.CreateNotification")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.CreateNotification", "error", err)
			SerializeLogMessage(Logger{
				Level:           LevelError,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.CreateNotification error: %s", err),
			})
		} else {
			a.logger.Info("<-- Notifications.CreateNotification")
			SerializeLogMessage(Logger{
				Level:           LevelInfo,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.CreateNotification"),
			})
		}
	}()

	return a.App.CreateNotification(ctx, cmd)
}

func (a Application) UpdateNotification(ctx context.Context, cmd commands.UpdateNotificationCommand) (notification *domain.Notification, err error) {
	a.logger.Info("--> Notifications.UpdateNotification")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.UpdateNotification", "error", err)
			SerializeLogMessage(Logger{
				Level:           LevelError,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.UpdateNotification error: %s", err),
			})
		} else {
			a.logger.Info("<-- Notifications.UpdateNotification")
			SerializeLogMessage(Logger{
				Level:           LevelInfo,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.UpdateNotification"),
			})
		}
	}()

	return a.App.UpdateNotification(ctx, cmd)
}

func (a Application) DeleteNotification(ctx context.Context, cmd commands.DeleteNotificationCommand) (err error) {
	a.logger.Info("--> Notifications.DeleteNotification")
	defer func() {
		if err != nil {
			a.logger.Error("<-- Notifications.DeleteNotification", "error", err)
			SerializeLogMessage(Logger{
				Level:           LevelError,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.DeleteNotification error: %s", err),
			})
		} else {
			a.logger.Info("<-- Notifications.DeleteNotification")
			SerializeLogMessage(Logger{
				Level:           LevelInfo,
				FromApplication: "NotificationService",
				DateTime:        time.Now().Format(time.RFC3339),
				Message:         fmt.Sprintf("<-- Notifications.DeleteNotification"),
			})
		}
	}()

	return a.App.DeleteNotification(ctx, cmd)
}
