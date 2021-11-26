package usecase

import (
	"github.com/jorgepuerta00/accountapi-master/model"
)

type ManageUseCaseInterface interface {
	Create(account *model.AccountData) (resultAccount *model.AccountData, err error)
	Delete(account *model.AccountData) error
}

type ListAccountUseCaseInterface interface {
	ListAccount() ([]model.AccountData, error)
	Find(id int) (*model.AccountData, error)
}

