package usecase

import (
	"github.com/SHU95/docker-env-go/domain"
)

type UserInteractor struct {
	userRepository UserRepository
}

func (interactor *UserInteractor) UserById(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByID(id)
	return
}

func (interactor *UserInteractor) Add(createUser domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Store(createUser)
	return
}

func (interactor *UserInteractor) Update(updateUser domain.User) (user domain.User, err error) {
	user, err = interactor.UserRespository.Update(updateUser)
	return
}
