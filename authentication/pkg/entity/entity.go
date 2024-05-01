package entity

type Entity struct {
	Id        ID
	IsActive  bool
	IsDeleted bool
}

func AlreadyEntity(id string, isActive, isDeleted bool) (*Entity, error) {

	idParsed, errParse := ParseID(id)

	if errParse != nil {
		if id == "" {
			return nil, ErrIDIsRequired
		}
		return nil, ErrInvalidID
	}

	entity := &Entity{
		Id:        idParsed,
		IsActive:  isActive,
		IsDeleted: isDeleted,
	}

	err := entity.validate()

	if err != nil {
		return nil, err
	}

	return entity, err
}

func NewEntity() (*Entity, error) {

	entity := &Entity{
		Id:        NewID(),
		IsActive:  true,
		IsDeleted: false,
	}

	err := entity.validate()

	if err != nil {
		return nil, err
	}

	return entity, err
}

func (e *Entity) validate() error {
	if e.Id.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := ParseID(e.Id.String()); err != nil {
		return ErrInvalidID
	}

	return nil
}
