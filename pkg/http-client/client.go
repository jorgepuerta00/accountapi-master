package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jorgepuerta00/accountapi-master/pkg/model"
	"github.com/sirupsen/logrus"
)

type ExternalSource interface {
	Create(model.Account) (model.Account, *http.Response, error)
	Delete(id string, version int) (bool, *http.Response, error)
	Get(id string) (model.Account, *http.Response, error)
	GetAll(model.PageParams) ([]model.Account, *http.Response, error)
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

func (c APIRecruitClient) Create(account model.Account) (model.Account, *http.Response, error) {
	body := &model.Result{Data: account}

	payload := new(bytes.Buffer)

	err := json.NewEncoder(payload).Encode(&body)
	if err != nil {
		return model.Account{}, nil, err
	}

	resp, err := c.httpClient.Post(c.baseURL, "application/json", payload)
	if err != nil {
		return model.Account{}, resp, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		createResult := &model.ErrorResult{}

		if err := json.NewDecoder(resp.Body).Decode(&createResult); err != nil {
			c.logger.Error("APIRecruitClient.Create", "error:", createResult.ErrorMessage)
			return model.Account{}, resp, err
		}

		return model.Account{}, resp, errors.New(createResult.ErrorMessage)
	}

	accountResult := &model.Result{}

	if err := json.NewDecoder(resp.Body).Decode(&accountResult); err != nil {
		c.logger.Error("APIRecruitClient.Create", "error:", err)
		return model.Account{}, resp, err
	}

	return accountResult.Data, resp, nil
}

func (c APIRecruitClient) Delete(id string, version int) (bool, *http.Response, error) {
	url := fmt.Sprintf("%s/%s?version=%d", c.baseURL, id, version)

	resp, err := c.customRequest(http.MethodDelete, url, nil)
	if err != nil {
		return false, resp, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		deleteResult := &model.ErrorResult{}

		if err := json.NewDecoder(resp.Body).Decode(&deleteResult); err != nil {
			return false, resp, err
		}

		return false, resp, errors.New(deleteResult.ErrorMessage)
	}

	return true, resp, nil
}

func (c APIRecruitClient) Get(id string) (model.Account, *http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, id)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return model.Account{}, resp, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		deleteResult := &model.ErrorResult{}

		if err := json.NewDecoder(resp.Body).Decode(&deleteResult); err != nil {
			c.logger.Error("APIRecruitClient.Get", "error:", deleteResult.ErrorMessage)
			return model.Account{}, resp, err
		}

		return model.Account{}, resp, errors.New(deleteResult.ErrorMessage)
	}

	getResponse := &model.Result{}

	if err := json.NewDecoder(resp.Body).Decode(&getResponse); err != nil {
		c.logger.Error("APIRecruitClient.Get", "error:", err)
		return model.Account{}, resp, err
	}

	return getResponse.Data, resp, nil
}

func (c APIRecruitClient) GetAll(pageParams model.PageParams) ([]model.Account, *http.Response, error) {

	url := solveParamsUrl(c.baseURL, pageParams)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return []model.Account{}, resp, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		deleteResult := &model.ErrorResult{}

		if err := json.NewDecoder(resp.Body).Decode(&deleteResult); err != nil {
			c.logger.Error("APIRecruitClient.Get", "error:", deleteResult.ErrorMessage)
			return []model.Account{}, resp, err
		}

		return []model.Account{}, resp, errors.New(deleteResult.ErrorMessage)
	}

	getallResponse := &model.ArrayResult{}

	if err := json.NewDecoder(resp.Body).Decode(&getallResponse); err != nil {
		c.logger.Error("APIRecruitClient.GetAll", "error:", err)
		return []model.Account{}, resp, err
	}

	return getallResponse.Data, resp, nil
}

func solveParamsUrl(baseURL string, pageParams model.PageParams) string {
	url, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}
	urlQuery := url.Query()

	if pageParams.Page >= 0 {
		urlQuery.Set("page[number]", strconv.Itoa(pageParams.Page))
	}
	if pageParams.Size >= 0 {
		urlQuery.Set("page[size]", strconv.Itoa(pageParams.Size))
	}

	url.RawQuery = urlQuery.Encode()

	return url.String()
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
