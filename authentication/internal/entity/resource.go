package entity

import "upse/authentication/pkg/entity"

type Resource struct {
	*entity.Entity
	Description string
	Key         string
}

func NewResourse(entity *entity.Entity, description string, key string) *Resource {
	return &Resource{
		Entity:      entity,
		Description: description,
		Key:         key,
	}
}
