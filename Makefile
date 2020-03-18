export GOPATH := $(shell pwd):$(GOPATH)

prog:
        go get -d ./...
        GOARCH=arm GOOS=linux go build
