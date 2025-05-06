package queries

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Logger/internal/domain"
)

type FindAllLogQueryHandler struct {
	log domain.LogRepository
}

func NewFindAllLogQuery(log domain.LogRepository) FindAllLogQueryHandler {
	return FindAllLogQueryHandler{log: log}
}
func (q FindAllLogQueryHandler) FindAll(ctx context.Context) ([]domain.Log, error) {
	logs, err := q.log.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return logs, nil
}
