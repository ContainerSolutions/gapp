build:
	@go build -o bin/app main.go

test:
	@go test -v .

run:
	@VERSION=0.1.0 go run main.go
