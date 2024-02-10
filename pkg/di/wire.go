//go:build wireinject
// +build wireinject

package di

import (
	"github.com/ajalck/spine/pkg/api"
	"github.com/ajalck/spine/pkg/api/handler"
	"github.com/ajalck/spine/pkg/config"
	"github.com/ajalck/spine/pkg/repository"
	"github.com/ajalck/spine/pkg/usecase"
	"github.com/ajalck/spine/pkg/db"
	"github.com/google/wire"
)

func InitializeAPI(c *config.Config) (*api.Server,error) {
	wire.Build(
		db.ConnectDB,
		//auth
		repository.NewAuthRepo,
		usecase.NewAuthUseCase,
		handler.NewAuthHandler,
		api.NewServerHTTP,
	)
	return &api.Server{},nil
}
