package database

import (
	"upse/authentication/internal/entity"
	pkg "upse/authentication/pkg/entity"
)

type (
	IPersonRepository interface {
		CreatePerson(person *entity.Person) error
		FindById(id pkg.ID) (*entity.Person, error)
		FindByName(name string) ([]entity.Person, error)
		FindAll(page, limit int, sort, by string) ([]entity.Person, error)
		Update(id pkg.ID, person *entity.Person) error
		Delete(id pkg.ID) error
	}

	IUserRepository interface {
		CreateUser(user *entity.User) error
	}
)
