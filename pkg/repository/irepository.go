package repository

import (
	"context"

	"github.com/jorgepuerta00/accountapi-master/pkg/model"
)

type AccountRepository interface {
	Create(context.Context, model.Account) (model.Account, error)
	Delete(context.Context, string) (bool, error)
	GetAll(context.Context) ([]model.Account, error)
	GetById(context.Context, string) (model.Account, error)
}
