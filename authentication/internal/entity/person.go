package entity

import (
	"errors"
	"upse/authentication/pkg/entity"
)

var (
	ErrIDIsRequired   = errors.New("ID is required")
	ErrInvalidID      = errors.New("invalid ID")
	ErrNameIsRequired = errors.New("name is required")
)

type Person struct {
	*entity.Entity
	Name string
}

func NewPerson(name string) (*Person, error) {
	person := &Person{
		Entity: entity.NewEntity(),
		Name:   name,
	}

	err := person.validate()

	if err != nil {
		return nil, err
	}

	return person, nil
}

func (p *Person) validate() error {

	if p.Id.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(p.Id.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	return nil
}
