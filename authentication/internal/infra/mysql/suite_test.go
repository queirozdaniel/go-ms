package mysql

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	ddlPersons := `
	CREATE TABLE persons (
		id BINARY(36) NOT NULL,
		name varchar(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		is_active TINYINT DEFAULT 1,
		is_deleted TINYINT DEFAULT 0,
		CONSTRAINT persons_pk PRIMARY KEY (id)
	);
	`
	_, err = db.Exec(ddlPersons)
	suite.NoError(err)

	ddlUsers := `
	CREATE TABLE users (
		id BINARY(36) NOT NULL,
		name varchar(255) NOT NULL,
		email varchar(255) NOT NULL,
		password varchar(125) NOT NULL,
		id_person BINARY(36) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		is_active TINYINT DEFAULT 1,
		is_deleted TINYINT DEFAULT 0,
		CONSTRAINT persons_pk PRIMARY KEY (id)
		CONSTRAINT users_persons_FK FOREIGN KEY (id_person) REFERENCES persons(id)
	);
	`
	_, err = db.Exec(ddlUsers)
	suite.NoError(err)

	suite.Db = db
}

func (suite *RepositoryTestSuite) TearDownTest() {
	_, err := suite.Db.Exec("delete from persons;")
	suite.NoError(err)

	_, err = suite.Db.Exec("delete from users;")
	suite.NoError(err)
}

func (suite *RepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}
