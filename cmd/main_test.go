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
	c, err := CreateClient()
	t.client = c
	assert.NoError(t.T(), err)
}

func (t *MainTestSuite) Test_Main_Create_Repository() {
	r, err := CreateRepository(t.client)
	t.repo = r
	assert.NoError(t.T(), err)
}

func (t *MainTestSuite) Test_Main_Create_Service() {
	s, err := CreateService(t.repo)
	t.service = s
	assert.NoError(t.T(), err)
}

func (t *MainTestSuite) Test_Main_Provisioning_Service() {
	s, err := CreateAccountService()
	t.service = s
	assert.NoError(t.T(), err)
}
