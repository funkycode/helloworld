FROM golang:1.14 AS builder
RUN go get github.com/funkycode/helloworld
WORKDIR /go/src/github.com/funkycode/helloworld
RUN CGO_ENABLED=0 GOOS=linux go build 

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/funkycode/helloworld/helloworld .
CMD ["./helloworld"]  
EXPOSE 8080
