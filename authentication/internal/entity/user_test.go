package entity

import (
	"testing"
	"upse/authentication/pkg/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	user, err := NewUser("daniel", "daniel@gmail.com", "123123")

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.NotEmpty(t, user.Name)
	assert.NotEmpty(t, user.Email)
	assert.NotEmpty(t, user.Password, "123123")
	assert.NotEqual(t, user.Password, "123123")
	assert.Equal(t, user.PersonId, user.Person.Id)
	assert.Equal(t, user.Name, user.Person.Name)
}

func TestNewUserInvalidName(t *testing.T) {
	user, err := NewUser("", "daniel@gmai.com", "123123")

	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, err, entity.ErrNameIsRequired)
}

func TestNewUserInvalidEmail(t *testing.T) {
	user, err := NewUser("dani", "", "123123")

	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, err, entity.ErrEmailIsRequired)
}

func TestNewUserInvalidPassword(t *testing.T) {
	user, err := NewUser("dani", "dam@gmail.com", "")

	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, err, entity.ErrPasswordIsRequired)
}

func TestUserValidPassword(t *testing.T) {
	user, err := NewUser("dani", "dam@gmail.com", "123123")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("123123"))
}

func TestUserInvalidPassword(t *testing.T) {
	user, err := NewUser("dani", "dam@gmail.com", "123123")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.False(t, user.ValidatePassword("11111"))
}
