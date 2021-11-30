package repository

import (
	httpclient "github.com/jorgepuerta00/accountapi-master/pkg/http-client"
	"github.com/jorgepuerta00/accountapi-master/pkg/model"

	"github.com/sirupsen/logrus"
)

type AccountRepository interface {
	Create(model.Account) (model.Account, error)
	Delete(id string, version int) (bool, error)
	GetAll(model.PageParams) ([]model.Account, error)
	GetById(id string) (model.Account, error)
}

type AccountRepo struct {
	externalHTTPClient httpclient.ExternalSource
	logger             logrus.FieldLogger
}

func NewAccountRepo(logger logrus.FieldLogger, externalHTTPClient httpclient.ExternalSource) *AccountRepo {
	return &AccountRepo{
		externalHTTPClient: externalHTTPClient,
		logger:             logger,
	}
}

func (mr *AccountRepo) Create(account model.Account) (model.Account, error) {
	accountResult, _, err := mr.externalHTTPClient.Create(account)
	if err != nil {
		mr.logger.Error("AccountRepo.Create", "fail calling http external api", "error", err)
		return model.Account{}, err
	}
	return accountResult, nil
}

func (mr *AccountRepo) Delete(id string, version int) (bool, error) {
	accountResult, _, err := mr.externalHTTPClient.Delete(id, version)
	if err != nil {
		mr.logger.Error("AccountRepo.Delete", "fail calling http external api", "error", err)
		return false, err
	}
	return accountResult, nil
}

func (mr *AccountRepo) GetAll(pageParams model.PageParams) ([]model.Account, error) {
	accountResult, _, err := mr.externalHTTPClient.GetAll(pageParams)
	if err != nil {
		mr.logger.Error("AccountRepo.GetAll", "fail calling http external api", "error", err)
		return []model.Account{}, err
	}
	return accountResult, nil
}

func (mr *AccountRepo) GetById(id string) (model.Account, error) {
	accountResult, _, err := mr.externalHTTPClient.Get(id)
	if err != nil {
		mr.logger.Error("AccountRepo.Get", "fail calling http external api", "error", err)
		return model.Account{}, err
	}
	return accountResult, nil
}
