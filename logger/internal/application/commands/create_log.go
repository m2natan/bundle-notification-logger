package commands

import (
	"context"
	"log/slog"

	"github.com/Kifiya-Financial-Technology/Logger/internal/domain"
)

type CreateLogCommand struct {
	Level           domain.LogLevel
	FromApplication string
	DateTime        string
	Message         string
}

type CreateLogHandler struct {
	log domain.LogRepository
}

func NewCreateNotificationHandler(log domain.LogRepository) CreateLogHandler {
	return CreateLogHandler{log: log}
}

func (h CreateLogHandler) CreateLog(ctx context.Context, cmd CreateLogCommand) (*domain.Log, error) {
	slog.Info("%s", cmd)
	log, err := domain.CreateLog(cmd.Level, cmd.FromApplication, cmd.DateTime, cmd.Message)
	if err != nil {
		return nil, err
	}

	err = h.log.Create(ctx, log)
	if err != nil {
		return nil, err
	}

	return log, nil
}
