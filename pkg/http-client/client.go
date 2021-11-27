package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jorgepuerta00/accountapi-master/pkg/model"
	"github.com/sirupsen/logrus"
)

type ExternalSource interface {
	Create(context.Context, model.Account) (model.Account, error)
	Delete(ctx context.Context, id string, version int) (bool, error)
	GetAll(context.Context) ([]model.Account, error)
	GetById(ctx context.Context, id string) (model.Account, error)
}

type body struct {
	Data interface{} `json:"data"`
}

type consumable interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (*http.Response, error)
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
}

type APIRecruitClient struct {
	baseURL    string
	httpClient consumable
	logger     logrus.FieldLogger
}

func NewAPIRecruitClient(logger logrus.FieldLogger, baseURL string) *APIRecruitClient {
	return &APIRecruitClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
		logger:     logger,
	}
}

func (c APIRecruitClient) Create(ctx context.Context, account model.Account) (model.Account, error) {
	body := body{Data: account}
	payload := new(bytes.Buffer)

	err := json.NewEncoder(payload).Encode(&body)
	if err != nil {
		return model.Account{}, err
	}

	resp, err := c.httpClient.Post(c.baseURL, "application/json", payload)
	if err != nil {
		return model.Account{}, err
	}

	defer resp.Body.Close()

	accountResponse := model.Account{}

	if err := json.NewDecoder(resp.Body).Decode(&accountResponse); err != nil {
		c.logger.Error("APIRecruitClient.Create", "error:", err)
		return model.Account{}, err
	}

	return accountResponse, nil
}

func (c APIRecruitClient) Delete(ctx context.Context, id string, version int) (bool, error) {
	url := fmt.Sprintf("%s/%s?version=%d", c.baseURL, id, version)

	resp, err := c.customRequest(http.MethodDelete, url, nil)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	accountResponse := model.Account{}

	if err := json.NewDecoder(resp.Body).Decode(&accountResponse); err != nil {
		c.logger.Error("APIRecruitClient.Create", "error:", err)
		return false, err
	}

	return true, nil
}

func (c APIRecruitClient) GetAll(ctx context.Context) ([]model.Account, error) {
	resp, err := c.httpClient.Get(c.baseURL)
	if err != nil {
		return []model.Account{}, err
	}

	defer resp.Body.Close()

	accountResponse := []model.Account{}

	if err := json.NewDecoder(resp.Body).Decode(&accountResponse); err != nil {
		c.logger.Error("APIRecruitClient.GetAll", "error:", err)
		return []model.Account{}, err
	}

	return accountResponse, nil
}

func (c APIRecruitClient) GetById(ctx context.Context, id string) (model.Account, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, id)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return model.Account{}, err
	}

	defer resp.Body.Close()

	accountResponse := model.Account{}

	if err := json.NewDecoder(resp.Body).Decode(&accountResponse); err != nil {
		c.logger.Error("APIRecruitClient.Get", "error:", err)
		return model.Account{}, err
	}

	return accountResponse, nil
}

func (c *APIRecruitClient) customRequest(method string, url string, body io.Reader) (*http.Response, error) {
	var req *http.Request
	var resp *http.Response
	var err error

	req, err = http.NewRequest(method, url, body)
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
