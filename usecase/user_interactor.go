package usecase

import "strategy-user/domain"

type UserInteractor struct {
	UserRepo UserRepo
}

func (interactor *UserInteractor) Add(user domain.UserModel) (err error) {
	_, err = interactor.UserRepo.Store(user)
	return 
}

func (interactor *UserInteractor) UserById(id int64) (user domain.UserModel, err error) {
	user, err = interactor.UserRepo.FindById(id)
	return
}


func (interactor *UserInteractor) Users(id int64) (users []domain.UserModel, err error) {
	users, err = interactor.UserRepo.FindAll()
	return
}