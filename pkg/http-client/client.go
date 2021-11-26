package httpclient

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type consumable interface {
	Get(string) (*http.Response, error)
}

func NewAPIRecruitClient(logger logrus.FieldLogger, baseURL string) *APIRecruitClient {
	return &APIRecruitClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
		logger:     logger,
	}
}

type APIRecruitClient struct {
	baseURL    string
	httpClient consumable
	logger     logrus.FieldLogger
}
