GO = go

.PHONY: all

all: api-gateway

api-gateway:
	GO111MODULE=on CC=gcc GOOS=linux CGO_ENABLED=0 $(GO) build -v -mod=vendor -o $@

clean:
	rm api-gateway
