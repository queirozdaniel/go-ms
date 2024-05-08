package entity

import (
	"upse/authentication/pkg/entity"

	"golang.org/x/crypto/bcrypt"
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

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	return user, err
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) ChangePassword(newPassword string) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	u.Password = string(hashPassword)

	return err
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
