package service

import (
	"os"
	"testing"

	"github.com/google/uuid"
	httpclient "github.com/jorgepuerta00/accountapi-master/pkg/http-client"
	"github.com/jorgepuerta00/accountapi-master/pkg/model"
	"github.com/jorgepuerta00/accountapi-master/pkg/repository"
	"github.com/jorgepuerta00/accountapi-master/pkg/seed"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(AccountServiceTestSuite))
}

type AccountServiceTestSuite struct {
	suite.Suite
	repo             *repository.AccountRepo
	service          *AccountService
	seed             *seed.Seeding
	testAccountsData []model.Account
}

func (t *AccountServiceTestSuite) SetupTest() {
	logger := logrus.New()
	baseUrl := os.Getenv("BASE_URL")
	client := httpclient.NewAPIRecruitClient(logger, baseUrl)
	repo := repository.NewAccountRepo(logger, client)

	t.service = NewAccountService(logger, repo)
	t.seed = seed.NewSeeding(logger, baseUrl)
	data, _ := t.seed.Seeding(5)
	t.testAccountsData = data
}

func (t *AccountServiceTestSuite) Test_Service_Create_Success() {
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

func (t *AccountServiceTestSuite) Test_Service_Create_Duplicated() {

	newAccount := model.Account{
		Type:           "accounts",
		ID:             t.testAccountsData[0].ID,
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
	assert.Error(t.T(), err)
	assert.Equal(t.T(), model.Account{}, result)
}

func (t *AccountServiceTestSuite) Test_Service_Delete_NotFound() {
	result, err := t.service.Delete(uuid.NewString(), 0)
	assert.Error(t.T(), err)
	assert.Equal(t.T(), false, result)
}

func (t *AccountServiceTestSuite) Test_Service_Delete_Invalid() {
	result, err := t.service.Delete("123", 0)
	assert.Error(t.T(), err)
	assert.Equal(t.T(), false, result)
}

func (t *AccountServiceTestSuite) Test_Service_Delete_wrongVersion() {
	result, err := t.service.Delete(t.testAccountsData[0].ID, 1)
	assert.Error(t.T(), err)
	assert.Equal(t.T(), false, result)
}

func (t *AccountServiceTestSuite) Test_Service_Delete_Sucess() {
	result, err := t.service.Delete(t.testAccountsData[0].ID, 0)
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), true, result)
}

func (t *AccountServiceTestSuite) Test_Service_Get_Sucess() {
	result, err := t.service.Fetch(t.testAccountsData[0].ID)
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), t.testAccountsData[0].ID, result.ID)
}

func (t *AccountServiceTestSuite) Test_Service_Get_Invalid() {
	result, err := t.service.Fetch("123")
	assert.Error(t.T(), err)
	assert.Equal(t.T(), model.Account{}, result)
}

func (t *AccountServiceTestSuite) Test_Service_Get_NotFound() {
	result, err := t.service.Fetch("ad27e265-9605-4b4b-a0e5-3003ea9cc4dd")
	assert.Error(t.T(), err)
	assert.Equal(t.T(), model.Account{}, result)
}

func (t *AccountServiceTestSuite) Test_Service_GetAll() {
	result, err := t.service.List(model.PageParams{})
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(result) > 0)
}

func (t *AccountServiceTestSuite) Test_Service_GetAll_Page() {
	result, err := t.service.List(model.PageParams{Page: 5})
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(result) == 0)
}

func (t *AccountServiceTestSuite) Test_Service_GetAll_Size() {
	result, err := t.service.List(model.PageParams{Size: 5})
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(result) > 0)
}

func (t *AccountServiceTestSuite) Test_Service_GetAll_PagginingParams() {
	result, err := t.service.List(model.PageParams{Page: 5, Size: 5})
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(result) > 0)
}
