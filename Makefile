export GOPATH := $(shell pwd):$(GOPATH)

prog:
		go get -d ./...
		GOARCH=arm GOOS=linux go build
		GOARCH=arm GOOS=linux go build -o tools/lc-tlscert tools/lc-tlscert.go
