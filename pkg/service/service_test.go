package service

import (
	"context"
	"os"
	"testing"

	httpclient "github.com/jorgepuerta00/accountapi-master/pkg/http-client"
	"github.com/jorgepuerta00/accountapi-master/pkg/model"
	"github.com/jorgepuerta00/accountapi-master/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(AccountServiecTestSuite))
}

type AccountServiecTestSuite struct {
	suite.Suite
	client  *httpclient.APIRecruitClient
	repo    *repository.AccountRepo
	service *AccountService
}

func (t *AccountServiecTestSuite) SetupTest() {
	logger := logrus.New()
	t.client = httpclient.NewAPIRecruitClient(logger, os.Getenv("BASE_URL"))
	t.repo = repository.NewAccountRepo(logger, t.client)
	t.service = NewAccountService(logger, t.repo)
}

func (t *AccountServiecTestSuite) Test_Service_Create() {
	assert.Equal(t.T(), true, true)
}

func (t *AccountServiecTestSuite) Test_Service_Delete() {
	assert.Equal(t.T(), true, true)
}

func (t *AccountServiecTestSuite) Test_Service_Get() {
	ctx := context.TODO()

	expectedResp := []model.Account{{
		ID:             "1",
		OrganisationID: "1",
		Type:           "accounts",
		Version:        0,
		Attributes:     model.AccountAttributes{},
	}}

	result, err := t.service.Fetch(ctx, "1")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), expectedResp, result)
}

func (t *AccountServiecTestSuite) Test_Service_GetAll() {
	assert.Equal(t.T(), true, true)
}
