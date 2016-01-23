export GO15VENDOREXPERIMENT=1

all: install

deps:
	go get github.com/Masterminds/glide
	glide --home $(HOME) up
	go get -d $$(glide nv)

install:
	go install -v $$(glide nv)

test:
	go test -v $$(glide nv)

.PHONY: all deps install test
