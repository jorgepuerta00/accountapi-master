package repository

import "accountapi-master/pkg/model"

func Create(account *model.Account) (resultAccount *model.Account, err error) {
	return nil, nil
}

func Delete(id string) (bool, error) {
	return true, nil
}

func GetAll() ([]model.Account, error) {
	return nil, nil
}

func Get(id string) (*model.Account, error) {
	return nil, nil
}
