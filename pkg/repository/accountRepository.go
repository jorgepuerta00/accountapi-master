package repository

import "accountapi-master/pkg/model"

type AccountRepository interface {
	Create(account *model.Account) (resultAccount *model.Account, err error)
	Delete(id string) (bool, error)
	GetAll() ([]model.Account, error)
	Get(id string) (*model.Account, error)
}
