#This repository is archived and will no longer receive updates.

GoREST
======

Learning how to write a RESTful app in Go - part of my [100 Days of Code challenge](https://github.com/clcollins/100-days-of-code/blob/master/log.md).

I've never written a RESTful app, and I don't know Go, so here I \*cough\* go...

How to Build and Run the App
----------------------------

1. Clone the repository
2. `bash build-me.sh && docker-compose up`

_Prerequesites:_ 

* docker
* docker-compose
* ...bash and git

Go Project Structure
--------------------

* build-me.sh: Builds the Dockerfile-builder image, copies the binary to ./pkg, and then builds the Dockerfile image, using the binary from before
* Dockerfile-builder: Instructions for building the Go binary from source using the Go base image
* Dockerfile: Instructions for building an image `FROM scratch` with the Go binary in it
* docker-compose.yml: Starts up a container running my go app
