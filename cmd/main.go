package main

import (
	"fmt"
	"os"

	httpclient "github.com/jorgepuerta00/accountapi-master/pkg/http-client"
	"github.com/jorgepuerta00/accountapi-master/pkg/repository"
	"github.com/jorgepuerta00/accountapi-master/pkg/service"
	"github.com/sirupsen/logrus"
)

var (
	baseUrl = os.Getenv("BASE_URL")
	logger  = logrus.New()
)

func main() {
	fmt.Println("client library form3 build")
}

func CreateClient() (*httpclient.APIRecruitClient, error) {
	return httpclient.NewAPIRecruitClient(logger, baseUrl), nil
}

func CreateRepository(client *httpclient.APIRecruitClient) (*repository.AccountRepo, error) {
	return repository.NewAccountRepo(logger, client), nil
}

func CreateService(repo *repository.AccountRepo) (*service.AccountService, error) {
	return service.NewAccountService(logger, repo), nil
}

func CreateAccountService() (*service.AccountService, error) {

	client, err := CreateClient()
	repository, err := CreateRepository(client)
	service, err := CreateService(repository)

	return service, err
}
