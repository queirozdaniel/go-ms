package entity

import "upse/authentication/pkg/entity"

type Role struct {
	*entity.Entity
	Description string
	Resources   []Resource
}

func NewRole(entity *entity.Entity, description string, resources []Resource) *Role {
	return &Role{
		Entity:      entity,
		Description: description,
		Resources:   resources,
	}
}
