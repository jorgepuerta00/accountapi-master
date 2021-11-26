package repository

import (
	"context"

	httpclient "github.com/jorgepuerta00/accountapi-master/pkg/http-client"
	"github.com/jorgepuerta00/accountapi-master/pkg/model"

	"github.com/sirupsen/logrus"
)

func NewAccountRepo(logger logrus.FieldLogger, externalHTTPClient httpclient.ExternalSource) *AccountRepo {
	return &AccountRepo{
		externalHTTPClient: externalHTTPClient,
		logger:             logger,
	}
}

type AccountRepo struct {
	externalHTTPClient httpclient.ExternalSource
	logger             logrus.FieldLogger
}

func (mr *AccountRepo) Create(ctx context.Context, account model.Account) (model.Account, error) {
	accountResult, err := mr.externalHTTPClient.Create(ctx, account)
	if err != nil {
		mr.logger.Error(
			"AccountRepo.Create",
			"fail calling http external api",
			"error",
			err,
		)
		return model.Account{}, err
	}
	return accountResult, nil
}

func (mr *AccountRepo) Delete(ctx context.Context, id string) (bool, error) {
	accountResult, err := mr.externalHTTPClient.Delete(ctx, id)
	if err != nil {
		mr.logger.Error(
			"AccountRepo.Delete",
			"fail calling http external api",
			"error",
			err,
		)
		return false, err
	}
	return accountResult, nil
}

func (mr *AccountRepo) GetAll(ctx context.Context) ([]model.Account, error) {
	accountResult, err := mr.externalHTTPClient.GetAll(ctx)
	if err != nil {
		mr.logger.Error(
			"AccountRepo.GetAll",
			"fail calling http external api",
			"error",
			err,
		)
		return []model.Account{}, err
	}
	return accountResult, nil
}

func (mr *AccountRepo) GetById(ctx context.Context, id string) (model.Account, error) {
	accountResult, err := mr.externalHTTPClient.GetById(ctx, id)
	if err != nil {
		mr.logger.Error(
			"AccountRepo.Get",
			"fail calling http external api",
			"error",
			err,
		)
		return model.Account{}, err
	}
	return accountResult, nil
}
