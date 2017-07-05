REVISION := $(shell git describe --always)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)
LDFLAGS	:= -ldflags="-X \"main.Revision=$(REVISION)\" -X \"main.BuildDate=${DATE}\""

.PHONY: build deps clean run

build:
	 go build -o bin/c2d $(LDFLAGS) cmd/c2d/*.go
deps:
	dep ensure
clean:
	rm -f bin/c2d
run:
	go run cmd/c2d/c2d.go -c example/config.toml
