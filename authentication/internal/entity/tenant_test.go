package entity

import (
	"testing"
	"upse/authentication/pkg/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewTenant(t *testing.T) {

	tenant, err := NewTenant("UPSE Digital LTDA", "Fake CNPJ")

	assert.NotNil(t, tenant)
	assert.Nil(t, err)
}

func TestNewTenantInvalidName(t *testing.T) {

	tenant, err := NewTenant("", "CNPJ")

	assert.NotNil(t, err)
	assert.Nil(t, tenant)
	assert.Equal(t, err, entity.ErrNameIsRequired)
}

func TestNewTenantInvalidExternalKey(t *testing.T) {

	tenant, err := NewTenant("Dev Daniel LTDA", "")

	assert.NotNil(t, err)
	assert.Nil(t, tenant)
	assert.Equal(t, err, entity.ErrExternalKeyIsRequired)
}
