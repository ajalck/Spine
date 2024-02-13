package interfaces

import (
	"github.com/ajalck/spine/pkg/api/handler/request"
	"github.com/ajalck/spine/pkg/api/handler/response"
)

type AuthUseCase interface {
	SignUp(data *request.User) (*response.User, error)
}
