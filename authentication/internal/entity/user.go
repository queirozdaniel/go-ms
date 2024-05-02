package entity

import (
	"upse/authentication/pkg/entity"
)

type User struct {
	*entity.Entity
	Name     string
	Email    string
	Password string
	PersonId entity.ID
	Person   Person
}

func NewUser(name, email, password string) (*User, error) {

	entity, err := entity.NewEntity()

	if err != nil {
		return nil, err
	}

	person, err := NewPerson(name)

	if err != nil {
		return nil, err
	}

	user := &User{
		Entity:   entity,
		Name:     name,
		Email:    email,
		Password: password,
		PersonId: person.Id,
		Person:   *person,
	}

	err = user.validate()

	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *User) validate() error {

	if u.Name == "" {
		return entity.ErrNameIsRequired
	}

	if u.Email == "" {
		return entity.ErrEmailIsRequired
	}

	if u.Password == "" {
		return entity.ErrPasswordIsRequired
	}

	return nil
}
