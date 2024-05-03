package mysql

import (
	"database/sql"
	"upse/authentication/internal/entity"
)

type PersonRepository struct {
	DB *sql.DB
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{
		DB: db,
	}
}

func (pr *PersonRepository) CreatePerson(person *entity.Person) error {
	stmt, err := pr.DB.Prepare("insert into person(id, name, created_at, updated_at, is_active, is_deleted) values (?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(person.Id, person.Name, person.CreatedAt, person.UpdatedAt, person.IsActive, person.IsDeleted)
	if err != nil {
		return err
	}

	return nil
}
