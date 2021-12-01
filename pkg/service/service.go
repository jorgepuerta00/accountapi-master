package service

import (
	"github.com/jorgepuerta00/accountapi-master/pkg/model"
	"github.com/jorgepuerta00/accountapi-master/pkg/repository"

	"github.com/sirupsen/logrus"
)

type Service interface {
	Create(model.Account) (resultAccount model.Account, err error)
	Delete(id string, version int) (bool, error)
	List(model.PageParams) ([]model.Account, error)
	Fetch(id string) (*model.Account, error)
}

type AccountService struct {
	repo   repository.AccountRepository
	logger logrus.FieldLogger
}

func NewAccountService(logger logrus.FieldLogger, repo repository.AccountRepository) *AccountService {
	return &AccountService{
		repo:   repo,
		logger: logger,
	}
}

func (ms *AccountService) Create(account model.Account) (model.Account, error) {
	accountResult, err := ms.repo.Create(account)
	if err != nil {
		return model.Account{}, err
	}
	return accountResult, nil
}

func (ms *AccountService) Delete(id string, version int) (bool, error) {
	accountResult, err := ms.repo.Delete(id, version)
	if err != nil {
		return accountResult, err
	}
	return accountResult, nil
}

func (ms *AccountService) List(pageParams model.PageParams) ([]model.Account, error) {
	accountResult, err := ms.repo.GetAll(pageParams)
	if err != nil {
		return []model.Account{}, err
	}
	return accountResult, nil
}

func (ms *AccountService) Fetch(id string) (model.Account, error) {
	accountResult, err := ms.repo.GetById(id)
	if err != nil {
		return model.Account{}, err
	}
	return accountResult, nil
}
