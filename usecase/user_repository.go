package usecase

import "github.com/SHU95/docker-env-go/domain"

type UserRespository interface {
	FindByID(id int) (domain.User, error)
	Store(domain.User) (domain.User, error)
	Update(domain.User) (domain.User, error)
}
