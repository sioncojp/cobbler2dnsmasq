REVISION := $(shell git describe --always)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)
LDFLAGS	:= -ldflags="-X \"main.Revision=$(REVISION)\" -X \"main.BuildDate=${DATE}\""

.PHONY: build-cross dist build deps clean run

build-cross:
	GOOS=linux GOARCH=amd64 go build -o bin/c2d-linux-amd64 $(LDFLAGS) cmd/c2d/*.go
	GOOS=darwin GOARCH=amd64 go build -o bin/c2d-darwin-amd64 $(LDFLAGS) cmd/c2d/*.go

dist: build-cross
	cd bin && tar zcvf c2d-linux-amd64.tar.gz c2d-linux-amd64 && rm -f c2d-linux-amd64
	cd bin && tar zcvf c2d-darwin-amd64.tar.gz c2d-darwin-amd64 && rm -f c2d-darwin-amd64

build:
	 go build -o bin/c2d $(LDFLAGS) cmd/c2d/*.go
deps:
	dep ensure
clean:
	rm -f bin/c2d
run:
	go run cmd/c2d/c2d.go -c example/config.toml
