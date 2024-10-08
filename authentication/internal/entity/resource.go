package entity

import "upse/authentication/pkg/entity"

type Resource struct {
	*entity.Entity
	Description string
	Key         string
}

func NewResourse(description string, key string) (*Resource, error) {

	entity, err := entity.NewEntity()

	if err != nil {
		return nil, err
	}

	resource := &Resource{
		Entity:      entity,
		Description: description,
		Key:         key,
	}

	err = resource.validade()

	if err != nil {
		return nil, err
	}

	return resource, err
}

func (r *Resource) validade() error {

	if r.Description == "" {
		return entity.ErrDescriptionIsRequired
	}

	if r.Key == "" {
		return entity.ErrKeyIsRequired
	}

	return nil
}
