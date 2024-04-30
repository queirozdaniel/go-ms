package entity

type Entity struct {
	Id        ID
	IsActive  bool
	IsDeleted bool
}

func NewEntityExists(id ID, isActive, isDeleted bool) *Entity {
	return &Entity{
		Id:        id,
		IsActive:  isActive,
		IsDeleted: isDeleted,
	}
}

func NewEntity() *Entity {
	return &Entity{
		Id:        NewID(),
		IsActive:  true,
		IsDeleted: false,
	}
}
