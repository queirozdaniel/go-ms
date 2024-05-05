package mysql

import (
	"database/sql"
	"testing"
	"upse/authentication/internal/entity"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type PersonRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(PersonRepositoryTestSuite))
}

func (suite *PersonRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	ddl := `
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
	_, err = db.Exec(ddl)

	suite.NoError(err)
	suite.Db = db
}

func (suite *PersonRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func (suite *PersonRepositoryTestSuite) TestCreateOnePerson() {
	person, err := entity.NewPerson("Daniel")

	suite.NoError(err)
	repo := NewPersonRepository(suite.Db)

	err = repo.CreatePerson(person)
	suite.NoError(err)

}
