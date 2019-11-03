FROM golang:1.13.3-buster

ENV GO111MODULE on

RUN apt-get update && apt-get install -y --no-install-recommends \
            autoconf=2.69-11 \
            automake=1:1.16.1-4 \
            libtool=2.4.6-9 \
            && rm -rf /var/lib/apt/lists/*

ENV PROTOBUF_VERSION 3.10.1
RUN wget -O /usr/local/src/protobuf.tar.gz "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protobuf-cpp-${PROTOBUF_VERSION}.tar.gz"
WORKDIR /usr/local/src
RUN tar xzf protobuf.tar.gz
WORKDIR /usr/local/src/protobuf-${PROTOBUF_VERSION}
RUN ./autogen.sh && ./configure && make && make install && ldconfig

ENV GO_PROTOBUF_VERSION 1.3.2
RUN git clone https://github.com/golang/protobuf /go/src/github.com/golang/protobuf
WORKDIR /go/src/github.com/golang/protobuf
RUN git checkout "v${GO_PROTOBUF_VERSION}"
RUN go install github.com/golang/protobuf/protoc-gen-go

ENV GRPC_GATEWAY_VERSION 1.11.3
RUN git clone https://github.com/grpc-ecosystem/grpc-gateway /go/src/github.com/grpc-ecosystem/grpc-gateway
WORKDIR /go/src/github.com/grpc-ecosystem/grpc-gateway
RUN git checkout "v${GRPC_GATEWAY_VERSION}"
RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

ENV GRPC_WEB_VERSION 1.0.6
RUN wget -O /usr/local/bin/protoc-gen-grpc-web "https://github.com/grpc/grpc-web/releases/download/${GRPC_WEB_VERSION}/protoc-gen-grpc-web-${GRPC_WEB_VERSION}-linux-x86_64"
RUN chmod +x /usr/local/bin/protoc-gen-grpc-web

ENV GO_PROTO_VALIDATORS_VERSION 0.2.0
RUN git clone https://github.com/mwitkow/go-proto-validators /go/src/github.com/mwitkow/go-proto-validators
WORKDIR /go/src/github.com/mwitkow/go-proto-validators
RUN git checkout "v${GO_PROTO_VALIDATORS_VERSION}"
RUN go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

ENV PROTOC_GEN_DOC_VERSION 1.3.0
RUN git clone https://github.com/pseudomuto/protoc-gen-doc /go/src/github.com/pseudomuto/protoc-gen-doc
WORKDIR /go/src/github.com/pseudomuto/protoc-gen-doc
RUN git checkout "v${PROTOC_GEN_DOC_VERSION}"
RUN go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

ENV PROTOTOOL_VERSION 1.9.0
RUN wget "https://github.com/uber/prototool/releases/download/v${PROTOTOOL_VERSION}/prototool-$(uname -s)-$(uname -m)" \
            -O /usr/local/bin/prototool && \
            chmod +x /usr/local/bin/prototool

WORKDIR /
