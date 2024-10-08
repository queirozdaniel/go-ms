package mysql

import (
	"database/sql"
	"errors"
	"upse/authentication/internal/entity"
	pkg "upse/authentication/pkg/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user *entity.User) error {
	stmt, err := ur.db.Prepare("insert into users (id, name, email, password, id_person, created_at, updated_at, is_active, is_deleted) values (?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Name, user.Email, user.Password, user.PersonId, user.CreatedAt, user.UpdatedAt, 1, 0)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) FindById(id pkg.ID) (*entity.User, error) {
	stmt, err := ur.db.Prepare(`
		SELECT id, name, email, created_at, updated_at, is_active, id_person from users u 
		where id = ? and is_deleted = 0
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var u entity.User
	var e pkg.Entity

	err = stmt.QueryRow(id).Scan(&e.Id, &u.Name, &u.Email, &e.CreatedAt, &e.UpdatedAt, &e.IsActive, &u.PersonId)
	if err != nil {
		return nil, err
	}

	u.Entity = &e

	return &u, nil
}

func (ur *UserRepository) FindByPersonId(id pkg.ID) (*entity.User, error) {
	stmt, err := ur.db.Prepare(`
		SELECT id, name, email, created_at, updated_at, is_active, id_person from users u 
		where id_person = ? and u.is_deleted = 0
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var u entity.User
	var e pkg.Entity

	err = stmt.QueryRow(id).Scan(&e.Id, &u.Name, &u.Email, &e.CreatedAt, &e.UpdatedAt, &e.IsActive, &u.PersonId)
	if err != nil {
		return nil, err
	}

	u.Entity = &e

	return &u, nil
}

func (ur *UserRepository) ChangePassword(id pkg.ID, oldPassword, newPassword string) error {
	stmt, err := ur.db.Prepare("SELECT password from users where id = ? and is_deleted = 0")

	if err != nil {
		return err
	}

	defer stmt.Close()

	var u entity.User

	err = stmt.QueryRow(id).Scan(&u.Password)
	if err != nil {
		return err
	}

	if !u.ValidatePassword(oldPassword) {
		return errors.New("oldpassword is wrong")
	}

	err = u.ChangePassword(newPassword)

	if err != nil {
		return err
	}

	stmt, err = ur.db.Prepare("UPDATE users SET password = ? where id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Password, id)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(id pkg.ID) error {
	_, err := ur.Exists(id)

	if err != nil {
		return err
	}

	stmt, err := ur.db.Prepare("update users set is_deleted = 1 where id = ?")

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

func (ur *UserRepository) Exists(id pkg.ID) (*entity.User, error) {
	stmt, err := ur.db.Prepare("select id from users where id = ? and is_deleted = 0")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var u entity.User
	var e pkg.Entity

	err = stmt.QueryRow(id).Scan(&e.Id)
	if err != nil {
		return nil, err
	}

	u.Entity = &e

	return &u, nil
}
