package entity

import (
	"testing"
	"upse/authentication/pkg/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewResource(t *testing.T) {
	resource, err := NewResourse("Acesso aos arquivos", "document::access")

	assert.NotNil(t, resource)
	assert.Nil(t, err)
	assert.NotEmpty(t, resource.Description)
	assert.Equal(t, resource.Description, "Acesso aos arquivos")
	assert.Equal(t, resource.Key, "document::access")
}

func TestNewResourceInvalidDescription(t *testing.T) {
	resource, err := NewResourse("", "document::downloading")

	assert.NotNil(t, err)
	assert.Nil(t, resource)
	assert.Equal(t, err, entity.ErrDescriptionIsRequired)
}

func TestNewResourceInvalidKey(t *testing.T) {
	resource, err := NewResourse("Download Arquivos", "")

	assert.NotNil(t, err)
	assert.Nil(t, resource)
	assert.Equal(t, err, entity.ErrKeyIsRequired)
}
