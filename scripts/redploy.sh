#!/bin/bash

docker stop helloworld &>/dev/null
docker rm -f helloworld &>/dev/null
docker build . -t helloworld
docker run -d --name helloworld -p8080:8080 helloworld
