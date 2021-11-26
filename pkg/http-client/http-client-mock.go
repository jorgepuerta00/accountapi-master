package httpclient

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type httpClientMock struct {
	mock.Mock
}

func (m *httpClientMock) Get(url string) (*http.Response, error) {
	args := m.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}
