package entity

import "upse/authentication/pkg/entity"

type Tenant struct {
	*entity.Entity
	Name string
}

func NewTenant(entity *entity.Entity, name string) *Tenant {
	return &Tenant{
		Entity: entity,
		Name:   name,
	}
}
