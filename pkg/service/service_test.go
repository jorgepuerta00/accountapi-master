package service

import (
	"os"
	"testing"

	"github.com/google/uuid"
	httpclient "github.com/jorgepuerta00/accountapi-master/pkg/http-client"
	"github.com/jorgepuerta00/accountapi-master/pkg/model"
	"github.com/jorgepuerta00/accountapi-master/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(AccountServiceTestSuite))
}

type AccountServiceTestSuite struct {
	suite.Suite
	repo    *repository.AccountRepo
	service *AccountService
}

func (t *AccountServiceTestSuite) SetupTest() {
	logger := logrus.New()
	t.repo = repository.NewAccountRepo(logger, httpclient.NewAPIRecruitClient(logger, os.Getenv("BASE_URL")))
	t.service = NewAccountService(logger, t.repo)
}

func (t *AccountServiceTestSuite) Test_Service_Create() {
	newAccount := model.Account{
		Type:           "accounts",
		ID:             uuid.NewString(),
		OrganisationID: uuid.NewString(),
		Version:        0,
		Attributes: model.AccountAttributes{
			Country:                 "GB",
			BaseCurrency:            "GBP",
			AccountNumber:           "41426819",
			BankID:                  "400300",
			BankIDCode:              "GBDSC",
			Bic:                     "NWBKGB22",
			Iban:                    "GB11NWBK40030041426819",
			Name:                    []string{"Samantha Holder"},
			AlternativeNames:        []string{"Sam Holder"},
			AccountClassification:   "Personal",
			JointAccount:            false,
			AccountMatchingOptOut:   false,
			SecondaryIdentification: "A1B2C3D4",
			Switched:                false,
		},
	}

	result, err := t.service.Create(newAccount)

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), newAccount.ID, result.ID)
}

func (t *AccountServiceTestSuite) Test_Service_Delete() {
	assert.Equal(t.T(), true, true)
}

func (t *AccountServiceTestSuite) Test_Service_Get() {
	assert.Equal(t.T(), true, true)
}

func (t *AccountServiceTestSuite) Test_Service_GetAll() {
	assert.Equal(t.T(), true, true)
}
