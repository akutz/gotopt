export GO15VENDOREXPERIMENT=1

all: install

deps:
	go get github.com/Masterminds/glide
	glide --home $(HOME) up
	go get -d $$(glide nv)
	go get -v github.com/axw/gocov/gocov
	go get -v github.com/mattn/goveralls
	go get -v golang.org/x/tools/cmd/cover

install:
	go install -v $$(glide nv)

test: cover.out
cover.out:
	go test -v $$(glide nv) -cover -coverprofile cover.out

cover: cover.out
	goveralls -coverprofile=cover.out

.PHONY: all deps install test cover
