# Name: Jorge Andres Puerta Fernandez

## Summary

The client library was developed using Go which consumes a Form3 API.

Methods exposed:

* Create - creates one account
* Delete - delete one account by ID
* List - get all accounts using pagination or not 
* Fetch - get one account by ID

### Build solution

run ```docker-compose up```

### Architecture

I've tried to implement a domain-oriented solution, I've split the project into layers, the first layer is http-client which works as my data source, then I have my repository, in short, the Repository pattern allows me to easily test my application with unit tests, at the same time I'm able to separate business from data, the last layer is service, where I'm able to have a collection of repositories and sub-services that builds together the business flow  

I must be honest, my DDD implementation is not complete and undoubtedly needs to be refactorized.

### Good Practises SOLID

The broad goal of the SOLID principles is to reduce dependencies so that I'm able to change one module without impacting others. Additionally, they’re intended to make designs easier to understand, maintain, and extend. Ultimately, using these design principles makes it easier for me to avoid issues and to build adaptive, effective, and agile software

I've add dependency inyection on all layers, basically, I pass a dependency to another object or structure. I did this due to it allows me the creation of dependencies outside the dependant object. This is useful as I can decouple dependency creation from the object being created. 

### Unit Testing and Testing Integration

I tried to use test-driven development practices, although not all methods were covered due to time

* coverage is upon 90%
* every method was tested
* happy and unhappy path
* tests organized 

### Usage

```go
baseUrl := os.Getenv("http://localhost:8080/v1/organisation/accounts")
client  := httpclient.NewAPIRecruitClient(logger, baseUrl)
repo    := repository.NewAccountRepo(logger, client)
service := NewAccountService(logger, repo)

result, err := service.Create(model.Account{...})
result, err  = service.Delete("ad27e265-9605-4b4b-a0e5-3003ea9cc4dd", 1)
result, err  = service.Fetch("ad27e265-9605-4b4b-a0e5-3003ea9cc4dd")
result, err  = service.List(model.PageParams{})
result, err  = service.List(model.PageParams{Page: 5, Size: 5})
result, err  = service.List(model.PageParams{Page: 5})
result, err  = service.List(model.PageParams{Size: 5})
```
