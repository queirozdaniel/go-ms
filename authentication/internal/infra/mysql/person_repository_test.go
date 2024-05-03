package mysql

import (
	"database/sql"
	"log"
	"testing"
	"upse/authentication/internal/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var person, _ = entity.NewPerson("Daniel Queiroz")

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreatePerson(t *testing.T) {
	db, mock := NewMock()

	repository := NewPersonRepository(db)

	query := "INSERT INTO person \\(id, name, created_at, updated_at, is_active, is_deleted\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(person.Id, person.Name, person.CreatedAt, person.UpdatedAt, person.IsActive, person.IsDeleted)

	// query := "INSERT INTO users \\(id, name, email, phone\\) VALUES \\(\\?, \\?, \\?, \\?\\)"

	err := repository.CreatePerson(person)

	assert.Nil(t, err)
}
