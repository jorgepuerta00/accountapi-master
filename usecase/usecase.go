package usecase

import (
	"github.com/jorgepuerta00/accountapi-master/model"
)

type ManageUseCaseInterface interface {
	Create(account *model.Account) (resultAccount *model.Account, err error)
	Delete(account *model.Account) error
}

type ListAccountUseCaseInterface interface {
	ListAccount() ([]model.Account, error)
	Find(id int) (*model.Account, error)
}

