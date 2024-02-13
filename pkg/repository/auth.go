package repository

import (
	"github.com/ajalck/spine/pkg/domain"
	authrepo "github.com/ajalck/spine/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) authrepo.AuthRepo {
	return &authRepo{db}
}

func (r *authRepo) CreateUser(user *domain.User) (*domain.User, error) {

	tx := r.db.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
