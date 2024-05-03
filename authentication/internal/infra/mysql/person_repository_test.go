package mysql

import (
	"upse/authentication/internal/entity"

	"github.com/stretchr/testify/mock"
)

type PersonRepositoryMock struct {
	mock.Mock
}

func (pr *PersonRepositoryMock) CreatePerson(person *entity.Person) error {
	args := pr.Called(person)
	return args.Error(0)
}
