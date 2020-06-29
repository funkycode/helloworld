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

`DO NOT forget to update badge in the head of this README to link to your build status`

### Triggers

Each commit will trigger build tests

Each tag will result in new image pushed as

```$DOCKER_REGISTRY/$GITHUB_USERNAME/$GITHUB_REPO/helloworld:$GIT_TAG```

## Manual builds

### Build binary

Following will generate `helloworld` binary in root directory

```
git https://github.com/funkycode/helloworld.git
cd helloworld
# optionally run tests
go test ./...
go build 
```

### Build docker
```
git https://github.com/funkycode/helloworld.git
cd helloworld
docker build . -t helloworld
```

### Run docker manually
```
docker run --name helloworld -d -p8080:8080 helloworld # docker port 8080 is binded to host port 8080
```

### Update running image

```
# pull latest changes 
git pull 
sh scripts/redeploy.sh
```

Script above will do following:

1. It will rebuild image as `helloworld:latest` from latest code
2. Try to stop previous instance (in case you have one running already) 
3. Start docker with name `helloworld` while binding `8080` port of docker to `8080` on host.