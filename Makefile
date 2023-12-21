run:
	go run main.go
test:
	go test -race -v ./...
lint:
	golangci-lint run --fix