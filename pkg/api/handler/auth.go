package handler

import (
	authhandler "github.com/ajalck/spine/pkg/api/handler/interfaces"
	authusecase "github.com/ajalck/spine/pkg/usecase/interfaces"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	useCase authusecase.AuthUseCase
}

func NewAuthHandler(authUseCase authusecase.AuthUseCase) authhandler.AuthHandler {
	return &authHandler{
		authUseCase,
	}
}

func (a *authHandler) SignUp(ctx *gin.Context) {

}
func (a *authHandler) SignIn(ctx *gin.Context) {

}
