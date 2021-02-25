package database

import (
	"github.com/SHU95/docker-env-go/domain"
)

type UserRepository struct {
	SqlHandler
}

func (userRepo *UserRepository) FindbyID(id int) (user domain.User, err error) {
	err = userRepo.Find(&user, id).Error
	if err != nil {
		return
	}
	return
}

//mysqlに作ったユーザーを格納
func (userRepo *UserRepository) Store(createUser domain.User) (user domain.User, err error) {
	err = userRepo.Create(&createUser).Error
	if err != nil {
		return
	}
	user = createUser
	return
}

func (userRepo *UserRepository) Update(updateUser domain.User) (user domain.User, err error) {
	err = userRepo.Save(&updateUser).Error
	if err != nil {
		return
	}
	user = updateUser
	return
}
