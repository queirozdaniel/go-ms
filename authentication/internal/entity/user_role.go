package entity

import (
	"time"
	"upse/authentication/pkg/entity"
)

type UserRole struct {
	UserId    entity.ID
	User      User
	RoleId    entity.ID
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive  bool
	IsDeleted bool
}

func NewUserRole(userId entity.ID, roleId entity.ID) (*UserRole, error) {
	return &UserRole{
		UserId:    userId,
		RoleId:    roleId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
		IsDeleted: false,
	}, nil
}
