local-setup:
	go mod download

local-build:
	@go build -o exec cmd/main.go

tests:
	@go test ./... -count=1 -cover

local-run: local-build
	@./exec
	
docker-build:
	@docker-compose build

docker-run:
	@docker-compose up

run-tests:
	docker-compose up --build --abort-on-container-exit client