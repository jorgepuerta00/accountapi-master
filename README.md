# Jorge Andres Puerta Fernandez

Level Golang Experience: ~1 year

## Summary

The client library was developed using Go which consumes a Form3 API.

Methods exposed:

* Create - create one account
* Delete - delete one account by ID
* List - get all accounts using pagination or not 
* Fetch - get one account by ID

### Build solution

run ```docker-compose up```

### Architecture

I've tried to implement a domain-oriented solution, I've split the project into layers, the first layer is http-client which works as my data source, then I have my repository, in short, the Repository pattern allows me to easily test my application with unit tests, at the same time I'm able to separate business from data, the last layer is service, where I'm able to have a collection of repositories and sub-services that builds together the business flow  

### Practises Used

It's important to write clean code, it's one of the core concepts of software development, we know there are many software design approaches, patterns, and practices to ensure an understandable, flexible, and maintainable code base.

I've implemented SOLID principles, in my view, SOLID helps to reduce dependencies so that I'm able to change one module without impacting others. Additionally, they’re intended to make designs easier to understand, maintain, and extend. Ultimately, using these design principles makes it easier for me to avoid issues and to build adaptive, effective, and agile software

For example, I've added dependency injection on all layers, basically, I pass a dependency to another object or structure. I did this due to it allows me the creation of dependencies outside the dependant object. This is useful as I can decouple dependency creation from the object being created. 

### Testing

I tried to use test-driven development practices, basically, this works creating unit test cases before writing functional code. It is an iterative approach that combines programming, the creation of unit tests, and refactoring, although not all methods were covered due to time

* coverage is upon 90%
* every method was tested
* happy and unhappy path
* tests organized 

### Usage

```go
service := CreateAccountService()

result, err := service.Create(model.Account{...})
result, err  = service.Delete(uuid.NewString(), idVersion)
result, err  = service.Fetch(uuid.NewString())
result, err  = service.List(model.PageParams{})
result, err  = service.List(model.PageParams{Page: 5, Size: 5})
result, err  = service.List(model.PageParams{Page: 5})
result, err  = service.List(model.PageParams{Size: 5})
```
