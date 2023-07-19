.DEFAULT_GOAL := build

.PHONY: clean build test vet run 

vet:
	go vet ./...

test:
	go test -v ./...

build: vet
	go build -o unused-file-remover -gcflags "-m" -v ./unused-file-remover.go

clean:
	go clean
	rm -f unused-file-remover


run:
	go run ./unused-file-remover.go -verbose -maxage=100h -cachePath=/tmp/removetest/

