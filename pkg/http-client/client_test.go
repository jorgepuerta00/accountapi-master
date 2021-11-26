package httpclient

import (
	"testing"

	"github.com/sirupsen/logrus"
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
	t.client = NewAPIRecruitClient(logger, "")
}
