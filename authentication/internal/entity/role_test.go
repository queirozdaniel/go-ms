package entity

import (
	"testing"
	"upse/authentication/pkg/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewRole(t *testing.T) {
	role, err := NewRole("Analista de Arquivos", nil)

	assert.NotNil(t, role)
	assert.Nil(t, err)
	assert.NotEmpty(t, role.Description)
}

func TestNewRoleInvalidDescription(t *testing.T) {
	role, err := NewRole("", nil)

	assert.NotNil(t, err)
	assert.Nil(t, role)
	assert.Equal(t, err, entity.ErrDescriptionIsRequired)
}

func TestNewRoleWithResources(t *testing.T) {

	download, _ := NewResourse("Download de Arquivo", "document::download")
	upload, _ := NewResourse("Upload de Arquivo", "document::download")

	resources := []Resource{*download, *upload}

	role, err := NewRole("Analista de arquivos", resources)

	assert.NotNil(t, role)
	assert.Nil(t, err)
	assert.NotEmpty(t, role.Resources)
	assert.Equal(t, len(role.Resources), 2)
}
