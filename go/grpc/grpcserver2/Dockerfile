FROM golang:1.14.7-buster

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/110y/lab

COPY go.mod go.mod
COPY go.sum go.sum
RUN GOPROXY='direct' go mod download

COPY . .
RUN go build -o /usr/bin/grpcserver2 ./go/grpc/grpcserver2

CMD ["/usr/bin/grpcserver2"]
