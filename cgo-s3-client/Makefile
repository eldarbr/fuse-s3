DISABLED_LINTERS="depguard,paralleltest,execinquery,gochecknoglobals"

all:

.PHONY: build
build:
	gcc c-convert/c-convert.c -o c-convert/c-convert.o -c
	go build -buildmode=c-shared -o libs3-client.so cgo-s3-client.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	go vet ./...
	golangci-lint run --enable-all --color=never --disable=$(DISABLED_LINTERS)

clean:
	rm -rf libs3-client.* c-convert/*.o
