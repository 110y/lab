FROM golang:1.13.3-alpine3.10 AS builder

ENV GO111MODULE on

COPY . /usr/local/src/repo
WORKDIR /usr/local/src/repo

RUN go build -o server .

FROM alpine:3.10
COPY --from=builder /usr/local/src/repo/server /usr/local/bin/server
ENTRYPOINT ["/usr/local/bin/server"]
