# Playing with github, golang and travisCI

[![Build Status](https://travis-ci.org/funkycode/helloworld.svg?branch=master)](https://travis-ci.org/funkycode/helloworld)

## Description

This is simple ping-pong demo server that listens to `/` on port `8080` written in golang.
It accepts POST requests with json with following syntax:
```
{"name":"YOUR_NAME"}
```
If request is correct the response will be in following format:
```
{"greeting":"Hello, YOUR_NAME"}
```
In case of error following json will be returned:
```
{"error":"ERROR_DESCRIPTION"}
```
This repo contains travis CI integration and in case there is tag and all tests passed it will upload docker image as GitHub package, where image tag is the tag of git.

#### Curl usage example:
```
# curl -X POST -d '{"name":"world"}' localhost:8080
{"greeting":"Hello, world"}
```


## Travis CI

### Setup
In order to utilize travis CI on your fork, connect TravisCI to your repository.
Set following env variables for that repo:
`DOCKER_USER` - your user name on github
`DOCKER_PASSWORD` - API token for github
`DOCKER_REGISTRY` - registry to push to, e.g. for `Packages` on github mine is set to `docker.pkg.github.com`

### Triggers

Each commit will trigger build tests

Each tag will result in new image pushed as 
```$DOCKER_REGISTRY/$GITHUB_USERNAME/$GITHUB_REPO/helloworld:$GIT_TAG```

## Manual builds

### Build binary
```
git https://github.com/funkycode/helloworld.git
cd helloworld
# optionally run tests
go test ./...
go build 
```
This will generate helloworld binary in root directory

### Build docker
```
git https://github.com/funkycode/helloworld.git
cd helloworld
# optionally run tests
docker build . -t helloworld
```
### Run docker
```
docker run --name helloworld -d -p8080:8080 helloworld # docker port 8080 is binded to host port 8080
```
Or alternatively use `scripts/redeploy.sh`. It will rebuild image as `helloworld:latest`, try to stop previous instance (in case you have one running already) and start docker with name `helloworld` while binding `8080` port of docker to `8080` on host.

After that you can generate request using curl or other tool of your desire.