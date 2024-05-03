package database

import (
	"upse/authentication/internal/entity"
)

type (
	IPersonRepository interface {
		CreatePerson(person *entity.Person) error
	}

	IUserRepository interface {
		CreateUser(user *entity.User) error
	}
)
