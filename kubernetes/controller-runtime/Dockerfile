FROM golang:1.13.5-alpine3.10 AS builder

ENV GO111MODULE=on
WORKDIR /go/src/github.com/110y/lab
COPY . .
RUN go build -o /crexp ./kubernetes/controller-runtime


FROM alpine:3.10.2

COPY --from=builder /crexp /crexp
RUN apk update
RUN apk add --no-cache ca-certificates

CMD ["/crexp"]
