package entity

type Entity struct {
	Id        ID
	IsActive  bool
	IsDeleted bool
}

func NewEntity(id ID, isActive, isDeleted bool) *Entity {
	return &Entity{
		Id:        id,
		IsActive:  isActive,
		IsDeleted: isDeleted,
	}
}

func NewEntityDefault() *Entity {
	return &Entity{
		Id:        NewID(),
		IsActive:  true,
		IsDeleted: false,
	}
}
