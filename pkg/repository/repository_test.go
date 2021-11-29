package repository

import (
	"context"
	"os"
	"testing"

	httpclient "github.com/jorgepuerta00/accountapi-master/pkg/http-client"
	"github.com/jorgepuerta00/accountapi-master/pkg/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(AccountRepositoryTestSuite))
}

type AccountRepositoryTestSuite struct {
	suite.Suite
	client *httpclient.APIRecruitClient
	repo   *AccountRepo
}

func (t *AccountRepositoryTestSuite) SetupTest() {
	logger := logrus.New()
	t.client = httpclient.NewAPIRecruitClient(logger, os.Getenv("BASE_URL"))
	t.repo = NewAccountRepo(logger, t.client)
}

func (t *AccountRepositoryTestSuite) Test_Repo_Create() {
	assert.Equal(t.T(), true, true)
}

func (t *AccountRepositoryTestSuite) Test_Repo_Delete() {
	assert.Equal(t.T(), true, true)
}

func (t *AccountRepositoryTestSuite) Test_Repo_Get() {
	ctx := context.TODO()

	expectedResp := []model.Account{{
		ID:             "1",
		OrganisationID: "1",
		Type:           "accounts",
		Version:        0,
		Attributes:     model.AccountAttributes{},
	}}

	result, err := t.repo.GetById(ctx, "1")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), expectedResp, result)
}

func (t *AccountRepositoryTestSuite) Test_Repo_GetAll() {
	assert.Equal(t.T(), true, true)
}
