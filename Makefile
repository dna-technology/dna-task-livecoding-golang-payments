test:
	go test -v ./...

run:
	go run ./cmd

build:
	go build -o=./bin/payments ./cmd