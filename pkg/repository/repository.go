package repository

import (
	"context"

	"github.com/jorgepuerta00/accountapi-master/pkg/model"
)

type AccountRepository interface {
	Create(context.Context, model.Account) (model.Account, error)
	Delete(ctx context.Context, id string, version int) (bool, error)
	GetAll(context.Context) ([]model.Account, error)
	GetById(ctx context.Context, id string) (model.Account, error)
}
