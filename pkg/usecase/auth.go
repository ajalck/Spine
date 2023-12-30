package usecase

import (
	authrepo "github.com/ajalck/spine/pkg/repository/interfaces"
	authusecase "github.com/ajalck/spine/pkg/usecase/interfaces"
)

type authUseCase struct {
	repo authrepo.AuthRepo
}

func NewAuthUseCase(repo authrepo.AuthRepo) authusecase.AuthUseCase {
	return &authUseCase{repo}
}

func (u *authUseCase) SignUp(){

	
}
