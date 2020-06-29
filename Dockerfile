FROM scratch
ADD helloworld /
ENTRYPOINT /helloworld
EXPOSE 8080
