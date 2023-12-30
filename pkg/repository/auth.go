package repository

import (
	authrepo "github.com/ajalck/spine/pkg/repository/interfaces"
)

type authRepo struct {
}

func NewAuthRepo() authrepo.AuthRepo {
	return &authRepo{}
}

func (r *authRepo) SignUp() {

}
