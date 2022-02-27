package app

import (
	"context"
	"github.com/core-go/health"
	s "github.com/core-go/health/sql"
	"github.com/core-go/sql"
	_ "github.com/lib/pq"

	"go-service/internal/handler"
	"go-service/internal/service"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	UserHandler   *handler.UserHandler
}

func NewApp(context context.Context, conf Config) (*ApplicationContext, error) {
	db, err := sql.OpenByConfig(conf.Sql)
	if err != nil {
		return nil, err
	}

	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(userService)

	sqlChecker := s.NewHealthChecker(db)
	healthHandler := health.NewHandler(sqlChecker)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		UserHandler:   userHandler,
	}, nil
}
