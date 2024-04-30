package entity

import "upse/authentication/pkg/entity"

type User struct {
	entity.Entity
	AccessCode string
	Password   string
}
