package mysql

import (
	"database/sql"
	"fmt"
	"upse/authentication/internal/entity"
	paging "upse/authentication/pkg/database"
	pkg "upse/authentication/pkg/entity"
)

type PersonRepository struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{
		db: db,
	}
}

func (pr *PersonRepository) CreatePerson(person *entity.Person) error {
	stmt, err := pr.db.Prepare("insert into persons (id, name, created_at, updated_at, is_active, is_deleted) values (?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(person.Id, person.Name, person.CreatedAt, person.UpdatedAt, 1, 0)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PersonRepository) FindById(id pkg.ID) (*entity.Person, error) {
	stmt, err := pr.db.Prepare("select id, name, created_at, updated_at, is_active from persons where id = ? and is_deleted = 0")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var p entity.Person
	var e pkg.Entity

	err = stmt.QueryRow(id).Scan(&e.Id, &p.Name, &e.CreatedAt, &e.UpdatedAt, &e.IsActive)
	if err != nil {
		return nil, err
	}

	p.Entity = &e

	return &p, nil
}

func (pr *PersonRepository) FindByName(name string) ([]entity.Person, error) {
	rows, err := pr.db.Query("select id, name, created_at, updated_at, is_active from persons where name like ? and is_deleted = 0", "%"+name+"%")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var persons []entity.Person

	for rows.Next() {
		var p entity.Person
		var e pkg.Entity
		err = rows.Scan(&e.Id, &p.Name, &e.CreatedAt, &e.UpdatedAt, &e.IsActive)
		if err != nil {
			return nil, err
		}
		p.Entity = &e
		persons = append(persons, p)
	}

	return persons, nil
}

func (pr *PersonRepository) FindAll(page, limit int, sort, by string) ([]entity.Person, error) {
	pagingQuery := paging.PagingBy(page, limit, sort, by)

	rows, err := pr.db.Query(fmt.Sprintf("select id, name, created_at, updated_at, is_active from persons where is_deleted = 0 %s", pagingQuery))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var persons []entity.Person

	for rows.Next() {
		var p entity.Person
		var e pkg.Entity
		err = rows.Scan(&e.Id, &p.Name, &e.CreatedAt, &e.UpdatedAt, &e.IsActive)
		if err != nil {
			return nil, err
		}
		p.Entity = &e
		persons = append(persons, p)
	}

	return persons, nil
}

func (pr *PersonRepository) Update(id pkg.ID, person *entity.Person) error {
	_, err := pr.Exists(id)

	if err != nil {
		return err
	}

	stmt, err := pr.db.Prepare("update persons set name = ? where id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(person.Name, id)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PersonRepository) Delete(id pkg.ID) error {
	_, err := pr.Exists(id)

	if err != nil {
		return err
	}

	stmt, err := pr.db.Prepare("update persons set is_deleted = 1 where id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PersonRepository) Exists(id pkg.ID) (*entity.Person, error) {
	stmt, err := pr.db.Prepare("select id from persons where id = ? and is_deleted = 0")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var p entity.Person
	var e pkg.Entity

	err = stmt.QueryRow(id).Scan(&e.Id)
	if err != nil {
		return nil, err
	}

	p.Entity = &e

	return &p, nil
}
