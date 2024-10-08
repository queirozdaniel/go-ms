package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEntity(t *testing.T) {
	entity, err := NewEntity()

	assert.Nil(t, err)
	assert.NotNil(t, entity.Id)
	assert.Equal(t, entity.IsActive, true)
	assert.Equal(t, entity.IsDeleted, false)
}

func TestNewEntityWithInvalidId(t *testing.T) {

	entity, err := AlreadyEntity("123123", true, false)

	assert.NotNil(t, err)
	assert.Nil(t, entity)
	assert.Equal(t, ErrInvalidID, err)

}

func TestNewEntityWithEmptyId(t *testing.T) {

	entity, err := AlreadyEntity("", true, false)

	assert.NotNil(t, err)
	assert.Nil(t, entity)
	assert.Equal(t, ErrIDIsRequired, err)

}
