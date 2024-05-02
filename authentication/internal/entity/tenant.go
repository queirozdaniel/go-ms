package entity

import (
	"upse/authentication/pkg/entity"
)

type Tenant struct {
	*entity.Entity
	Name        string
	ExternalKey string
}

func NewTenant(name, externalKey string) (*Tenant, error) {

	entity, err := entity.NewEntity()

	if err != nil {
		return nil, err
	}

	tenant := &Tenant{
		Entity:      entity,
		Name:        name,
		ExternalKey: externalKey,
	}

	err = tenant.validate()

	if err != nil {
		return nil, err
	}

	return tenant, err
}

func (t *Tenant) validate() error {

	if t.Name == "" {
		return entity.ErrNameIsRequired
	}

	if t.ExternalKey == "" {
		return entity.ErrExternalKeyIsRequired
	}

	return nil
}
