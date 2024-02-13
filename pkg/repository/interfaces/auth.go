package interfaces

import "github.com/ajalck/spine/pkg/domain"

type AuthRepo interface {
	CreateUser(user *domain.User) (*domain.User, error)
}
