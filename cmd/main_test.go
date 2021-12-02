package main

import (
	"testing"

	httpclient "github.com/jorgepuerta00/accountapi-master/pkg/http-client"
	"github.com/jorgepuerta00/accountapi-master/pkg/repository"
	"github.com/jorgepuerta00/accountapi-master/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}

type MainTestSuite struct {
	suite.Suite
	client  *httpclient.APIRecruitClient
	repo    *repository.AccountRepo
	service *service.AccountService
}

func (t *MainTestSuite) SetupTest() {

}

func (t *MainTestSuite) Test_Main_Create_Client() {
	c := CreateClient()
	t.client = c
	assert.NotNil(t.T(), c)
	assert.NotEmpty(t.T(), c)
}

func (t *MainTestSuite) Test_Main_Create_Repository() {
	r := CreateRepository(t.client)
	t.repo = r
	assert.NotNil(t.T(), r)
	assert.NotEmpty(t.T(), r)
}

func (t *MainTestSuite) Test_Main_Create_Service() {
	s := CreateService(t.repo)
	t.service = s
	assert.NotNil(t.T(), s)
	assert.NotEmpty(t.T(), s)
}

func (t *MainTestSuite) Test_Main_Provisioning_Service() {
	s := CreateAccountService()
	t.service = s
	assert.NotNil(t.T(), s)
	assert.NotEmpty(t.T(), s)
}
