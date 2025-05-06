package domain

import "context"

type LogRepository interface {
	Create(ctx context.Context, log *Log) error
	FindAll(ctx context.Context) ([]Log, error)
}
