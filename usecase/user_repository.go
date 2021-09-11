package usecase

import "github.com/halhal23/strategy-user/domain"

type UserRepo interface {
	Store(domain.UserModel) (int64, error)
	FindById(int64) (domain.UserModel, error)
	FindAll() ([]domain.UserModel, error)
}