package service

import (
	"accountapi-master/pkg/model"
	"accountapi-master/pkg/repository"

	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateAccount(account *model.Account) (resultAccount *model.Account, err error)
	DeleteAccount(id string) (bool, error)
	ListAccounts() ([]model.Account, error)
	FindAccount(id string) (*model.Account, error)
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

func (ms *AccountService) CreateAccount(account *model.Account) (resultAccount *model.Account, err error) {
	accountResult, err := ms.repo.Create(account)
	if err != nil {
		return &model.Account{}, err
	}
	return accountResult, nil
}

func (ms *AccountService) DeleteAccount(id string) (bool, error) {
	accountResult, err := ms.repo.Delete(id)
	if err != nil {
		return accountResult, err
	}
	return accountResult, nil
}

func (ms *AccountService) ListAccounts() ([]model.Account, error) {
	accountResult, err := ms.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return accountResult, nil
}

func (ms *AccountService) FindAccount(id string) (*model.Account, error) {
	accountResult, err := ms.repo.Get(id)
	if err != nil {
		return &model.Account{}, err
	}
	return accountResult, nil
}
