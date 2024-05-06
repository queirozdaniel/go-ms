package mysql

import (
	"database/sql"
	"fmt"
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
	_, err := suite.Db.Exec("delete from persons;")
	suite.NoError(err)
}

func (suite *PersonRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func (suite *PersonRepositoryTestSuite) TestCreatePerson() {
	person, err := entity.NewPerson("Daniel")

	suite.NoError(err)
	repo := NewPersonRepository(suite.Db)

	err = repo.CreatePerson(person)
	suite.NoError(err)
}

func (suite *PersonRepositoryTestSuite) TestGetPersonById() {
	person, err := entity.NewPerson("Daniel")

	suite.NoError(err)
	repo := NewPersonRepository(suite.Db)

	err = repo.CreatePerson(person)
	suite.NoError(err)

	p, err := repo.FindById(person.Id)

	suite.NoError(err)

	suite.Equal(p.Id, person.Id)
	suite.Equal(p.Name, person.Name)
	suite.Equal(p.CreatedAt, person.CreatedAt)
	suite.Equal(p.UpdatedAt, person.UpdatedAt)
	suite.Equal(p.IsActive, person.IsActive)
	suite.Equal(p.IsDeleted, person.IsDeleted)
}

func (suite *PersonRepositoryTestSuite) TestGetPersonByName() {
	pJones, _ := entity.NewPerson("Jones da Silva")
	pLuis, _ := entity.NewPerson("Luis Otavio")

	repo := NewPersonRepository(suite.Db)

	err := repo.CreatePerson(pJones)
	suite.NoError(err)

	err = repo.CreatePerson(pLuis)
	suite.NoError(err)

	persons, err := repo.FindByName(pJones.Name)

	suite.NoError(err)

	suite.Equal(1, len(persons))

	p := persons[0]
	suite.Equal(p.Id, pJones.Id)
	suite.Equal(p.Name, pJones.Name)
	suite.Equal(p.CreatedAt, pJones.CreatedAt)
	suite.Equal(p.UpdatedAt, pJones.UpdatedAt)
	suite.Equal(p.IsActive, pJones.IsActive)
	suite.Equal(p.IsDeleted, pJones.IsDeleted)
}

func (suite *PersonRepositoryTestSuite) TestGetAllPersons() {

	repo := NewPersonRepository(suite.Db)

	for i := range 7 {
		person, _ := entity.NewPerson(fmt.Sprintf("Daniel da Silva %d", i))
		err := repo.CreatePerson(person)
		suite.NoError(err)
	}

	for i := range 7 {
		person, _ := entity.NewPerson(fmt.Sprintf("John Snow %d", i))
		err := repo.CreatePerson(person)
		suite.NoError(err)
	}

	persons, err := repo.FindAll(1, 10, "", "")

	suite.NoError(err)
	suite.Equal(10, len(persons))

	persons, _ = repo.FindAll(1, 30, "", "")

	suite.Equal(14, len(persons))
}

func (suite *PersonRepositoryTestSuite) TestUpdatePersonName() {
	person, err := entity.NewPerson("Miles Moralles")

	suite.NoError(err)
	repo := NewPersonRepository(suite.Db)

	err = repo.CreatePerson(person)
	suite.NoError(err)

	err = repo.Update(person.Id, &entity.Person{Name: "Peter Parker"})
	suite.NoError(err)

	person, _ = repo.FindById(person.Id)

	suite.Equal(person.Name, "Peter Parker")
}

func (suite *PersonRepositoryTestSuite) TestDeletePerson() {
	pPeter, _ := entity.NewPerson("Peter Parker")
	pPeter2, _ := entity.NewPerson("Peter Parker 2")

	repo := NewPersonRepository(suite.Db)

	err := repo.CreatePerson(pPeter)
	suite.NoError(err)

	err = repo.CreatePerson(pPeter2)
	suite.NoError(err)

	persons, _ := repo.FindByName("Peter")
	suite.Equal(2, len(persons))

	err = repo.Delete(pPeter2.Id)
	suite.NoError(err)

	persons, _ = repo.FindByName("Peter")
	suite.Equal(1, len(persons))
}
