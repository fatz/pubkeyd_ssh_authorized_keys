EXECUTABLE := pubkeyd_ssh_authorized_keys
VERSION = $(shell git describe --tags)
GITHASH = $(shell git rev-parse HEAD)
DATETIME= $(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

all: test build

build:
	mkdir -p out
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o out/$(EXECUTABLE)-linux
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o out/$(EXECUTABLE).exe
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o out/$(EXECUTABLE)-darwin

test:
	go test ./onelogingh
	# go test ./cmd
