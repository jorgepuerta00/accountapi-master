package usecase

import (
	"accountapi-master/pkg/model"
)

type ManageAccountUseCaseInterface interface {
	Create(account *model.Account) (resultAccount *model.Account, err error)
	Delete(account *model.Account) error
}

type ListAccountUseCaseInterface interface {
	ListAccounts() ([]model.Account, error)
	Find(id int) (*model.Account, error)
}
