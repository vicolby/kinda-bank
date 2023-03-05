build:
	@go build -o bin/kinda-bank

run: build
	@bin/kinda-bank

test:
	@go test -v ./...
