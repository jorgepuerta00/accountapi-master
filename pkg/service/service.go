package service

import (
	"context"

	"github.com/jorgepuerta00/accountapi-master/pkg/model"
	"github.com/jorgepuerta00/accountapi-master/pkg/repository"

	"github.com/sirupsen/logrus"
)

type Service interface {
	Create(context.Context, model.Account) (resultAccount model.Account, err error)
	Delete(ctx context.Context, id string, version int) (bool, error)
	List(context.Context) ([]model.Account, error)
	Fetch(ctx context.Context, id string) (*model.Account, error)
}

func NewAccountService(logger logrus.FieldLogger, repo repository.AccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

type AccountService struct {
	repo   repository.AccountRepository
	logger logrus.FieldLogger
}

func (ms *AccountService) Create(ctx context.Context, account model.Account) (model.Account, error) {
	accountResult, err := ms.repo.Create(ctx, account)
	if err != nil {
		return model.Account{}, err
	}
	return accountResult, nil
}

func (ms *AccountService) Delete(ctx context.Context, id string, version int) (bool, error) {
	accountResult, err := ms.repo.Delete(ctx, id, version)
	if err != nil {
		return accountResult, err
	}
	return accountResult, nil
}

func (ms *AccountService) List(ctx context.Context) ([]model.Account, error) {
	accountResult, err := ms.repo.GetAll(ctx)
	if err != nil {
		return []model.Account{}, err
	}
	return accountResult, nil
}

func (ms *AccountService) Fetch(ctx context.Context, id string) (model.Account, error) {
	accountResult, err := ms.repo.GetById(ctx, id)
	if err != nil {
		return model.Account{}, err
	}
	return accountResult, nil
}
