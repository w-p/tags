.DEFAULT_GOAL := build

branch := $(shell git symbolic-ref HEAD | cut -d'/' -f3)

all: clean build install
	@:

clean:
	go clean
	-@rm ./bin/tags-$(branch)

build:
	go build -o ./bin/tags-$(branch)

install:
	go install
