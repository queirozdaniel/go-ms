package entity

import (
	"upse/authentication/pkg/entity"
)

type Person struct {
	*entity.Entity
	Name string
}

func NewPerson(name string) (*Person, error) {
	entity, err := entity.NewEntity()

	if err != nil {
		return nil, err
	}

	person := &Person{
		Entity: entity,
		Name:   name,
	}

	err = person.validate()

	if err != nil {
		return nil, err
	}

	return person, nil
}

func (p *Person) validate() error {

	if p.Name == "" {
		return entity.ErrNameIsRequired
	}

	return nil
}
