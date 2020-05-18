FROM golang:1.14.3-buster AS builder

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/110y/lab
COPY . .
RUN go build -o authority ./envoy/authority

# runtime image
FROM gcr.io/distroless/base:3c213222937de49881c57c476e64138a7809dc54

COPY --from=builder /go/src/github.com/110y/lab/authority /usr/local/bin/authority

EXPOSE 8080

CMD ["/usr/local/bin/authority"]
