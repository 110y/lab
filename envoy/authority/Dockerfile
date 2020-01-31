FROM golang:1.13.7-alpine3.10 AS builder

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/110y/lab
COPY . .
RUN go build -o authority ./envoy/authority

# runtime image
FROM alpine:3.10.2

COPY --from=builder /go/src/github.com/110y/lab/authority /usr/local/bin/authority
RUN apk update
RUN apk add --no-cache ca-certificates

EXPOSE 8080

CMD ["/usr/local/bin/authority"]
