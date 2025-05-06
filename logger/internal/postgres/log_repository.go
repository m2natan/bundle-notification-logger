package postgres

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Logger/internal/domain"
	"gorm.io/gorm"
)

type LogRepository struct {
	db *gorm.DB
}

var _ domain.LogRepository = (*LogRepository)(nil)

func NewLogRepository(db *gorm.DB) *LogRepository {
	return &LogRepository{db: db}
}

// Create implements domain.LogRepository.
func (l *LogRepository) Create(ctx context.Context, log *domain.Log) error {
	return l.db.WithContext(ctx).Create(log).Error
}

// FindAll implements domain.LogRepository.
func (l *LogRepository) FindAll(ctx context.Context) ([]domain.Log, error) {
	var logs []domain.Log
	err := l.db.WithContext(ctx).Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}
