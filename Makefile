lint:
	go fmt ./...

build:
	go build -o ./bin/bt ./cmd/bt/main.go

run:
	go run ./cmd/bt/main.go

install: build
	cp ./bin/bt ~/.local/bin/bt
