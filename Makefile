all: install

install:
	@go mod tidy

start:
	@go run main.go
