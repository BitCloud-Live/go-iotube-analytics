#!make
version := $(shell git describe --abbrev=0 --tags)

all: build

# The `validate` target checks for errors and inconsistencies in 
# our specification of an API. This target can check if we're 
# referencing inexistent definitions and gives us hints to where
# to fix problems with our API in a static manner.
validate:
	swagger validate ./pkg/openapi/swagger.yml

# The `gen` target depends on the `validate` target as
# it will only succesfully generate the code if the specification
# is valid.
# 
# Here we're specifying some flags:
# --target              the base directory for generating the files;
# --spec                path to the swagger specification;
# --exclude-main        generates only the library code and not a 
#                       sample CLI application;
# --name                the name of the application.
gen: validate 
	swagger generate server \
		--target=./pkg/openapi/swagger \
		--spec=./pkg/openapi/swagger.yml \
		--exclude-main \
		--name=polydefi

build:
	go build -o build/api cmd/main.go

clean:
	rm -rf build/*

run:    build
	export $$(cat .env | grep -v ^\# | xargs) && \
		./build/api

docker-build:
	docker build -t hub.docker.io/polystation/polydefi-api:$(version) . 

docker-push:
	docker push hub.docker.io/polystation/polydefi-api:$(version)



update-deps: download tidy vendor

download:
	go mod download

tidy:
	go mod tidy -v

vendor:
	go mod vendor
