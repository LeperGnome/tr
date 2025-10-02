.PHONY: test


lint:
	go fmt ./...

test:
	go test -v ./...

build:
ifeq ($(OS),Windows_NT)
	go build -o ./bin/bt.exe ./cmd/bt/main.go
else
	go build -o ./bin/bt ./cmd/bt/main.go
endif

run:
	go run ./cmd/bt/main.go

install: build
ifneq ($(OS),Windows_NT)
	cp ./bin/bt ~/.local/bin/bt
endif
