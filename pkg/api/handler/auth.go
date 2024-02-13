package handler

import (
	"net/http"

	authhandler "github.com/ajalck/spine/pkg/api/handler/interfaces"
	"github.com/ajalck/spine/pkg/api/handler/request"
	"github.com/ajalck/spine/pkg/api/handler/response"
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
	var body *request.User

	if err := ctx.ShouldBindJSON(&body); err != nil {
		res := response.ErrorResponse(http.StatusBadRequest, "Faile to bind inputs", err, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	user, err := a.useCase.SignUp(body)
	if err != nil {
		res := response.ErrorResponse(http.StatusConflict, "Failed to sign up", err, nil)
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
	}

	res := response.SuccessResponse(http.StatusCreated, "Sign up successfull", user)
	ctx.JSON(http.StatusCreated, res)

}
func (a *authHandler) SignIn(ctx *gin.Context) {

}
