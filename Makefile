test:
	go test -v ./...

run:
	go run ./cmd/api

build:
	go build -o=./bin/payments ./cmd/api