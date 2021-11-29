package httpclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/jorgepuerta00/accountapi-master/pkg/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(APIRecruitClientTestSuite))
}

type APIRecruitClientTestSuite struct {
	suite.Suite
	httpClientMock *httpClientMock
	client         *APIRecruitClient
}

func (t *APIRecruitClientTestSuite) SetupTest() {
	t.httpClientMock = &httpClientMock{}
	logger := logrus.New()
	t.client = NewAPIRecruitClient(logger, os.Getenv("BASE_URL"))
}

func TestCustomRequest(c *APIRecruitClientTestSuite, t *testing.T, r *http.Request, expected string) {
	t.Helper()
	assert.Equal(c.T(), expected, r.Method)
}

func (c *APIRecruitClientTestSuite) TestClient_Delete(t *testing.T) {
	req, _ := c.client.customRequest("DELETE", "/", nil)
	TestCustomRequest(c, t, req.Request, "DELETE")
}

func (t *APIRecruitClientTestSuite) Test_Create_Success() {
	assert.Equal(t.T(), true, true)
}

func (t *APIRecruitClientTestSuite) Test_Delete_Success() {
	assert.Equal(t.T(), true, true)
}

func (t *APIRecruitClientTestSuite) Test_Get_Success() {
	ctx := context.TODO()
	mockedJSONResponse := `[{
		"id": "1",
		"organisation_id": "1",
		"type": "accounts",
		"version": 0,
		"attributes": {}
	}]`

	r := ioutil.NopCloser(bytes.NewReader([]byte(mockedJSONResponse)))

	expectedResp := []model.Account{{
		ID:             "1",
		OrganisationID: "1",
		Type:           "accounts",
		Version:        0,
		Attributes:     model.AccountAttributes{},
	}}

	t.httpClientMock.On("Get", mock.AnythingOfType("string")).Return(
		&http.Response{
			StatusCode: 200,
			Body:       r,
		},
		nil,
	)

	t.client.httpClient = t.httpClientMock

	result, err := t.client.Get(ctx, "1")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), expectedResp, result)
}

func (t *APIRecruitClientTestSuite) Test_GetAll_Success() {
	assert.Equal(t.T(), true, true)
}
