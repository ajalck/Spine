package repository

import (
	authrepo "github.com/ajalck/spine/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) authrepo.AuthRepo {
	return &authRepo{db}
}

func (r *authRepo) SignUp() {

}
