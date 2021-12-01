package repository

import (
	"os"
	"testing"

	"github.com/google/uuid"
	httpclient "github.com/jorgepuerta00/accountapi-master/pkg/http-client"
	"github.com/jorgepuerta00/accountapi-master/pkg/model"
	"github.com/jorgepuerta00/accountapi-master/pkg/seed"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AccountRepositoryTestSuite))
}

type AccountRepositoryTestSuite struct {
	suite.Suite
	repo             *AccountRepo
	seed             *seed.Seeding
	testAccountsData []model.Account
}

func (t *AccountRepositoryTestSuite) SetupTest() {
	logger := logrus.New()
	baseUrl := os.Getenv("BASE_URL")
	t.repo = NewAccountRepo(logger, httpclient.NewAPIRecruitClient(logger, baseUrl))
	t.seed = seed.NewSeeding(logger, baseUrl)
	data, _ := t.seed.Seeding(5)
	t.testAccountsData = data
}

func (t *AccountRepositoryTestSuite) Test_Repo_Create_Success() {
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

	result, err := t.repo.Create(newAccount)
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), newAccount.ID, result.ID)
}

func (t *AccountRepositoryTestSuite) Test_Repo_Create_Duplicated() {

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

	result, err := t.repo.Create(newAccount)
	assert.Error(t.T(), err)
	assert.Equal(t.T(), model.Account{}, result)
}

func (t *AccountRepositoryTestSuite) Test_Repo_Delete_NotFound() {
	result, err := t.repo.Delete(uuid.NewString(), 0)
	assert.Error(t.T(), err)
	assert.Equal(t.T(), false, result)
}

func (t *AccountRepositoryTestSuite) Test_Repo_Delete_Invalid() {
	result, err := t.repo.Delete("123", 0)
	assert.Error(t.T(), err)
	assert.Equal(t.T(), false, result)
}

func (t *AccountRepositoryTestSuite) Test_Repo_Delete_wrongVersion() {
	result, err := t.repo.Delete(t.testAccountsData[0].ID, 1)
	assert.Error(t.T(), err)
	assert.Equal(t.T(), false, result)
}

func (t *AccountRepositoryTestSuite) Test_Repo_Delete_Sucess() {
	result, err := t.repo.Delete(t.testAccountsData[0].ID, 0)
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), true, result)
}

func (t *AccountRepositoryTestSuite) Test_Repo_Get_Sucess() {
	result, err := t.repo.GetById(t.testAccountsData[0].ID)
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), t.testAccountsData[0].ID, result.ID)
}

func (t *AccountRepositoryTestSuite) Test_Repo_Get_Invalid() {
	result, err := t.repo.GetById("123")
	assert.Error(t.T(), err)
	assert.Equal(t.T(), model.Account{}, result)
}

func (t *AccountRepositoryTestSuite) Test_Repo_Get_NotFound() {
	result, err := t.repo.GetById("ad27e265-9605-4b4b-a0e5-3003ea9cc4dd")
	assert.Error(t.T(), err)
	assert.Equal(t.T(), model.Account{}, result)
}

func (t *AccountRepositoryTestSuite) Test_Repo_GetAll() {
	result, err := t.repo.GetAll(model.PageParams{})
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(result) > 0)
}

func (t *AccountRepositoryTestSuite) Test_Repo_GetAll_Page() {
	result, err := t.repo.GetAll(model.PageParams{Page: 5})
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(result) == 0)
}

func (t *AccountRepositoryTestSuite) Test_Repo_GetAll_Size() {
	result, err := t.repo.GetAll(model.PageParams{Size: 5})
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(result) > 0)
}

func (t *AccountRepositoryTestSuite) Test_Repo_GetAll_PagginingParams() {
	result, err := t.repo.GetAll(model.PageParams{Page: 5, Size: 5})
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(result) > 0)
}
