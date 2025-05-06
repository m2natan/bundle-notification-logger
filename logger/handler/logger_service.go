package handler

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Logger/internal/application"
	"github.com/Kifiya-Financial-Technology/Logger/internal/application/commands"
	"github.com/Kifiya-Financial-Technology/Logger/internal/converter"
	"github.com/Kifiya-Financial-Technology/Logger/loggerpb"
	"google.golang.org/grpc"
)

type LoggerServiceHandler struct {
	app application.App
	loggerpb.UnimplementedLoggerServiceServer
}

func NewServer(app application.App, grpcServer *grpc.Server) {
	logger := &LoggerServiceHandler{app: app}
	loggerpb.RegisterLoggerServiceServer(grpcServer, logger)
}

func (l *LoggerServiceHandler) CreateLog(ctx context.Context, req *loggerpb.CreateLogRequest) (*loggerpb.Log, error) {
	cmd := commands.CreateLogCommand{
		FromApplication: req.FromApplication,
		Message:         req.Message,
		Level:           converter.ConvertLogLevelToDomain(req.Status),
		DateTime:        req.DateTime,
	}
	log, err := l.app.CreateLog(ctx, cmd)
	if err != nil {
		return nil, err
	}
	return converter.ConvertLoggerToPb(log), nil
}

func (l *LoggerServiceHandler) FindAllLogs(ctx context.Context, req *loggerpb.FindAllLogsRequest) (*loggerpb.FindAllLogsResponse, error) {
	logs, err := l.app.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var pbLogs []*loggerpb.Log
	for _, log := range logs {
		pbLogs = append(pbLogs, converter.ConvertLoggerToPb(&log))
	}

	return &loggerpb.FindAllLogsResponse{
		Logs: pbLogs,
	}, nil
}
