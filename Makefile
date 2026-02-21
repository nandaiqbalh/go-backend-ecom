build:
	@go build -o bin/go-backend-ecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-backend-ecom