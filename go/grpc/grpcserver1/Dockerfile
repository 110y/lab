FROM golang:1.15.2-buster

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/110y/lab/go/grpc/grpcserver1

COPY go.mod go.mod
COPY go.sum go.sum

COPY . .
RUN go build -o /usr/bin/grpcserver1 .

CMD ["/usr/bin/grpcserver1"]
