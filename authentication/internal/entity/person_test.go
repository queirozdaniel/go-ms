package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNormalPerson(t *testing.T) {

	person, err := NewPerson("Daniel Queiroz")

	assert.Nil(t, err)
	assert.NotNil(t, person.Id)
	assert.NotEmpty(t, person.Name)
	assert.Equal(t, "Daniel Queiroz", person.Name)
	assert.Equal(t, true, person.IsActive)
	assert.Equal(t, false, person.IsDeleted)
}

func TestNewPersonInvalid(t *testing.T) {
	person, err := NewPerson("")

	assert.Nil(t, person)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}
