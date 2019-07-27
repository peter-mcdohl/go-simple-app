# Dockerize go-simple-app

Build the docker image for [go-simple-app](https://github.com/kevinturnip/go-simple-app) and then run it on docker container

## Getting Started

* Build binary file `CGO_ENABLED=0 go build -o app .`
* Build docker image `docker build -t go-simple-app .`

## Installation

Run docker container `docker container run -d -p 8080:8080 go-simple-app`
