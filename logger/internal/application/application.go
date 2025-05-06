package application

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Logger/internal/application/commands"
	"github.com/Kifiya-Financial-Technology/Logger/internal/application/queries"
	"github.com/Kifiya-Financial-Technology/Logger/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}

	Commands interface {
		CreateLog(ctx context.Context, cmd commands.CreateLogCommand) (*domain.Log, error)
	}
	Queries interface {
		FindAll(ctx context.Context) ([]domain.Log, error)
	}
	Application struct {
		appCommands
		appQueries
	}

	appCommands struct {
		commands.CreateLogHandler
	}
	appQueries struct {
		queries.FindAllLogQueryHandler
	}
)

var _ App = (*Application)(nil)

func New(
	log domain.LogRepository,
) *Application {
	return &Application{
		appCommands: appCommands{
			commands.NewCreateNotificationHandler(log),
		},
		appQueries: appQueries{
			queries.NewFindAllLogQuery(log),
		},
	}
}
