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

func CreateClient() *httpclient.APIRecruitClient {
	return httpclient.NewAPIRecruitClient(logger, baseUrl)
}

func CreateRepository(client *httpclient.APIRecruitClient) *repository.AccountRepo {
	return repository.NewAccountRepo(logger, client)
}

func CreateService(repo *repository.AccountRepo) *service.AccountService {
	return service.NewAccountService(logger, repo)
}

func CreateAccountService() *service.AccountService {
	client := CreateClient()
	repository := CreateRepository(client)
	service := CreateService(repository)

	return service
}
