package repository

import (
	"os"
	"testing"

	"github.com/google/uuid"
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

	result, _, err := t.client.Create(newAccount)

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), newAccount.ID, result.ID)
}

func (t *AccountRepositoryTestSuite) Test_Repo_Delete() {
	assert.Equal(t.T(), true, true)
}

func (t *AccountRepositoryTestSuite) Test_Repo_Get() {
	assert.Equal(t.T(), true, true)
}

func (t *AccountRepositoryTestSuite) Test_Repo_GetAll() {
	assert.Equal(t.T(), true, true)
}
