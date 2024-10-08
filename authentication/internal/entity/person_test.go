package entity

import (
	"testing"
	"upse/authentication/pkg/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewNormalPerson(t *testing.T) {

	person, err := NewPerson("Daniel Queiroz")

	assert.Nil(t, err)
	assert.NotNil(t, person.Id)
	assert.NotEmpty(t, person.Name)
	assert.Equal(t, "Daniel Queiroz", person.Name)
}

func TestNewPersonInvalid(t *testing.T) {
	person, err := NewPerson("")

	assert.Nil(t, person)
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrNameIsRequired, err)
}
