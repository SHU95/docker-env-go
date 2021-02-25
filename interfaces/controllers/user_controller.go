package controllers

import (
	"strconv"

	"github.com/SHU95/docker-env-go/domain"
	"github.com/SHU95/docker-env-go/interfaces/database"
	"github.com/SHU95/docker-env-go/usecase"
)

type UserController struct {
	interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) CreateUser(context Context) (err error) {
	createUser := domain.User{}
	context.Bind(&createUser)
	user, err := controller.interactor.Add(createUser)

	if err != nil {
		context.JSON(500, NewError(err))
		return
	}
	context.JSON(201, user)
	return
}

func (controller *UserController) GetUser(context Context) (err error) {
	id, _ := strconv.Atoi(context.Param("id"))
	user, err := controller.interactor.UserById(id)

	if err != nil {
		context.JSON(500, NewError(err))
		return
	}
	context.JSON(200, user)
	return
}

func (controller *UserController) UpdateUser(context Context) (err error) {
	id, _ := strconv.Atoi(context.Param("id"))
	updateUser := domain.User{ID: id}
	context.Bind(&updateUser)

	user, err := controller.interactor.Update(updateUser)
}
