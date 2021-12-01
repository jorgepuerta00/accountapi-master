package httpclient

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jorgepuerta00/accountapi-master/pkg/model"
	"github.com/jorgepuerta00/accountapi-master/pkg/seed"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(APIRecruitClientTestSuite))
}

type APIRecruitClientTestSuite struct {
	suite.Suite
	httpClientMock   *httpClientMock
	client           *APIRecruitClient
	seed             *seed.Seeding
	testAccountsData []model.Account
}

func (t *APIRecruitClientTestSuite) SetupTest() {
	t.httpClientMock = &httpClientMock{}
	logger := logrus.New()
	baseUrl := os.Getenv("BASE_URL")
	t.client = NewAPIRecruitClient(logger, baseUrl)
	t.seed = seed.NewSeeding(logger, baseUrl)
	data, _ := t.seed.Seeding(5)
	t.testAccountsData = data
}

func (t *APIRecruitClientTestSuite) Test_Create_Success() {

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

	result, resp, err := t.client.Create(newAccount)
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), newAccount.ID, result.ID)
	assert.Equal(t.T(), http.StatusCreated, resp.StatusCode)
}

func (t *APIRecruitClientTestSuite) Test_Create_Duplicated() {

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

	result, resp, err := t.client.Create(newAccount)
	assert.Error(t.T(), err)
	assert.Equal(t.T(), model.Account{}, result)
	assert.Equal(t.T(), http.StatusConflict, resp.StatusCode)
}

func (t *APIRecruitClientTestSuite) Test_Delete_NotFound() {
	result, resp, err := t.client.Delete(uuid.NewString(), 0)
	assert.Error(t.T(), err)
	assert.Equal(t.T(), false, result)
	assert.Equal(t.T(), http.StatusNotFound, resp.StatusCode)
}

func (t *APIRecruitClientTestSuite) Test_Delete_Invalid() {
	result, resp, err := t.client.Delete("123", 0)
	assert.Error(t.T(), err)
	assert.Equal(t.T(), false, result)
	assert.Equal(t.T(), http.StatusBadRequest, resp.StatusCode)
}

func (t *APIRecruitClientTestSuite) Test_Delete_wrongVersion() {
	result, resp, err := t.client.Delete(t.testAccountsData[0].ID, 1)
	assert.Error(t.T(), err)
	assert.Equal(t.T(), false, result)
	assert.Equal(t.T(), http.StatusConflict, resp.StatusCode)
}

func (t *APIRecruitClientTestSuite) Test_Delete_Sucess() {
	result, resp, err := t.client.Delete(t.testAccountsData[0].ID, 0)
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), true, result)
	assert.Equal(t.T(), http.StatusNoContent, resp.StatusCode)
}

func (t *APIRecruitClientTestSuite) Test_Get_Sucess() {
	result, resp, err := t.client.Get(t.testAccountsData[0].ID)
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), t.testAccountsData[0].ID, result.ID)
	assert.Equal(t.T(), http.StatusOK, resp.StatusCode)
}

func (t *APIRecruitClientTestSuite) Test_Get_Invalid() {
	result, resp, err := t.client.Get("123")
	assert.Error(t.T(), err)
	assert.Equal(t.T(), model.Account{}, result)
	assert.Equal(t.T(), http.StatusBadRequest, resp.StatusCode)
}

func (t *APIRecruitClientTestSuite) Test_Get_NotFound() {
	result, resp, err := t.client.Get("ad27e265-9605-4b4b-a0e5-3003ea9cc4dd")
	assert.Error(t.T(), err)
	assert.Equal(t.T(), model.Account{}, result)
	assert.Equal(t.T(), http.StatusNotFound, resp.StatusCode)
}

func (t *APIRecruitClientTestSuite) Test_GetAll() {
	result, resp, err := t.client.GetAll(model.PageParams{})
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(result) > 0)
	assert.Equal(t.T(), http.StatusOK, resp.StatusCode)
}

func (t *APIRecruitClientTestSuite) Test_GetAll_Page() {
	result, resp, err := t.client.GetAll(model.PageParams{Page: 5})
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(result) == 0)
	assert.Equal(t.T(), http.StatusOK, resp.StatusCode)
}

func (t *APIRecruitClientTestSuite) Test_GetAll_Size() {
	result, resp, err := t.client.GetAll(model.PageParams{Size: 5})
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(result) > 0)
	assert.Equal(t.T(), http.StatusOK, resp.StatusCode)
}

func (t *APIRecruitClientTestSuite) Test_GetAll_PagginingParams() {
	result, resp, err := t.client.GetAll(model.PageParams{Page: 5, Size: 5})
	assert.NoError(t.T(), err)
	assert.True(t.T(), len(result) > 0)
	assert.Equal(t.T(), http.StatusOK, resp.StatusCode)
}
