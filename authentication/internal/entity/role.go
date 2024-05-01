package entity

import "upse/authentication/pkg/entity"

type Role struct {
	*entity.Entity
	Description string
	Resources   []Resource
}

func NewRole(description string, resources []Resource) (*Role, error) {

	entity, err := entity.NewEntity()

	if err != nil {
		return nil, err
	}

	role := &Role{
		Entity:      entity,
		Description: description,
		Resources:   resources,
	}

	err = role.validate()

	if err != nil {
		return nil, err
	}

	return role, err
}

func (r *Role) validate() error {

	if r.Description == "" {
		return entity.ErrDescriptionIsRequired
	}

	return nil
}
