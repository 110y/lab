FROM debian:10.1

RUN apt-get update -y && apt-get install -y python cmake default-jre git-core

WORKDIR /usr/local/src

RUN git clone https://github.com/emscripten-core/emsdk.git
WORKDIR /usr/local/src/emsdk
RUN ./emsdk install 1.39.1
RUN ./emsdk activate 1.39.1

WORKDIR /usr/local/src
RUN git clone https://github.com/emscripten-core/emscripten.git

WORKDIR /usr/local/src/workspace
