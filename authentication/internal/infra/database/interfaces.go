package database

import (
	"upse/authentication/internal/entity"
	pkg "upse/authentication/pkg/entity"
)

type (
	IPersonRepository interface {
		CreatePerson(person *entity.Person) error
		FindById(id pkg.ID) (*entity.Person, error)
		FindByName(name string) ([]entity.Person, error)
		FindAll(page, limit int, sort, by string) ([]entity.Person, error)
		Update(id pkg.ID, person *entity.Person) error
		Delete(id pkg.ID) error
	}

	IUserRepository interface {
		CreateUser(user *entity.User) error
		FindById(id pkg.ID) (*entity.User, error)
		FindByPersonId(id pkg.ID) (*entity.User, error)
		ChangePassword(id pkg.ID, oldPassword, newPassword string) error
		Delete(id pkg.ID) error
	}

	IRole interface {
		CreateRole(role *entity.Role) error
		CreateRoleWithResources(role *entity.Role, resources []*entity.Resource) error
		FindAll(page, limit int, sort, by string) ([]entity.Role, error)
		FindById(id pkg.ID) (*entity.Role, error)
		Delete(id pkg.ID) error
	}

	IResource interface {
		CreateResource(resource *entity.Resource) error
		AddResourceInRole(resourceId pkg.ID, roleId pkg.ID) error
		FindById(id pkg.ID) (*entity.Resource, error)
		FindAllByRoleId(roleId pkg.ID) ([]entity.Resource, error)
		Delete(id pkg.ID) error
	}
)
