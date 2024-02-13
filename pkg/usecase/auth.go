package usecase

import (
	"github.com/ajalck/spine/pkg/api/handler/request"
	"github.com/ajalck/spine/pkg/api/handler/response"
	"github.com/ajalck/spine/pkg/domain"
	authrepo "github.com/ajalck/spine/pkg/repository/interfaces"
	authusecase "github.com/ajalck/spine/pkg/usecase/interfaces"
	"github.com/ajalck/spine/pkg/utils"
)

type authUseCase struct {
	repo authrepo.AuthRepo
}

func NewAuthUseCase(repo authrepo.AuthRepo) authusecase.AuthUseCase {
	return &authUseCase{repo}
}

func (u *authUseCase) SignUp(data *request.User) (*response.User, error) {

	hashedPassword, err := utils.GenerateHashFromPassword(data.Password)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		FullName: data.FullName,
		Email:    data.Email,
		Password: hashedPassword,
	}
	user, err = u.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	resUser := &response.User{
		FullName: user.FullName,
		Email:    user.Email,
	}
	return resUser, nil
}
