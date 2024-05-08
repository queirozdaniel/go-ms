package mysql

import "upse/authentication/internal/entity"

func (suite *RepositoryTestSuite) TestCreateUser() {

	user, err := entity.NewUser("Daniel", "daniel@gmail.com", "112233")
	suite.NoError(err)

	person, err := entity.NewPerson(user.Name)
	suite.NoError(err)

	repoPerson := NewPersonRepository(suite.Db)
	err = repoPerson.CreatePerson(person)
	suite.NoError(err)

	repoUser := NewUserRepository(suite.Db)
	err = repoUser.CreateUser(user)
	suite.NoError(err)

}

func (suite *RepositoryTestSuite) TestFindCreatedUser() {
	user, err := entity.NewUser("Daniel", "daniel@gmail.com", "112233")
	suite.NoError(err)

	person, err := entity.NewPerson(user.Name)
	suite.NoError(err)

	repoPerson := NewPersonRepository(suite.Db)
	err = repoPerson.CreatePerson(person)
	suite.NoError(err)

	repoUser := NewUserRepository(suite.Db)
	err = repoUser.CreateUser(user)
	suite.NoError(err)

	userById, err := repoUser.FindById(user.Id)
	suite.NoError(err)

	userByPersonId, err := repoUser.FindByPersonId(user.Person.Id)
	suite.NoError(err)

	suite.Equal(user.Name, userById.Name)
	suite.Equal(user.Name, userByPersonId.Name)
	suite.Equal(user.PersonId, userByPersonId.PersonId)
	suite.Equal(user.PersonId, userByPersonId.PersonId)
}

func (suite *RepositoryTestSuite) TestChangeUserPassword() {
	user, err := entity.NewUser("Daniel", "daniel@gmail.com", "112233")
	suite.NoError(err)

	person, err := entity.NewPerson(user.Name)
	suite.NoError(err)

	repoPerson := NewPersonRepository(suite.Db)
	err = repoPerson.CreatePerson(person)
	suite.NoError(err)

	repoUser := NewUserRepository(suite.Db)
	err = repoUser.CreateUser(user)
	suite.NoError(err)

	err = repoUser.ChangePassword(user.Id, "112233", "111222")
	suite.NoError(err)
}

func (suite *RepositoryTestSuite) TestDeleteUser() {
	user, err := entity.NewUser("Daniel", "daniel@gmail.com", "112233")
	suite.NoError(err)

	person, err := entity.NewPerson(user.Name)
	suite.NoError(err)

	repoPerson := NewPersonRepository(suite.Db)
	err = repoPerson.CreatePerson(person)
	suite.NoError(err)

	repoUser := NewUserRepository(suite.Db)
	err = repoUser.CreateUser(user)
	suite.NoError(err)

	err = repoUser.Delete(user.Id)
	suite.NoError(err)

	_, err = repoUser.FindById(user.Id)
	suite.Error(err)

}
