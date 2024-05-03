package database

import (
	"upse/authentication/internal/entity"
	pkg "upse/authentication/pkg/entity"
)

type (
	IPersonRepository interface {
		CreatePerson(person *entity.Person) (pkg.ID, error)
	}

	IUserRepository interface {
		CreateUser(user *entity.User) (pkg.ID, error)
	}
)
