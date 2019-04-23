GO = go

.PHONY: all

all: api-gateway

api-gateway:
	GO111MODULE=on CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 $(GO) build -v -mod=vendor -o $@

clean:
	rm api-gateway