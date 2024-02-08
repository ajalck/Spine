//+build wireinject

package di

import (
	"github.com/ajalck/spine/pkg/api"
	"github.com/ajalck/spine/pkg/api/handler"
	"github.com/ajalck/spine/pkg/repository"
	"github.com/ajalck/spine/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI() *api.Server {
	wire.Build(
		//auth
		repository.NewAuthRepo,
		usecase.NewAuthUseCase,
		handler.NewAuthHandler,
		api.NewServerHTTP,
	)
	return &api.Server{}
}
