FROM envoyproxy/envoy-build-ubuntu:d06dad145694f1a7a02b5c6d0c75b32f753db2dd

RUN apt-get purge --auto-remove -y cmake

ENV CMAKE_MINOR_VERSION=3.15
ENV CMAKE_PATCH_VERSION=5

RUN wget https://cmake.org/files/v$CMAKE_MINOR_VERSION/cmake-$CMAKE_MINOR_VERSION.$CMAKE_PATCH_VERSION.tar.gz -O /usr/local/src/cmake.tar.gz
RUN tar xvzf /usr/local/src/cmake.tar.gz && \
            cd cmake-$CMAKE_MINOR_VERSION.$CMAKE_PATCH_VERSION && \
            ./bootstrap && \
            make && \
            make install && \
            rm -rf /usr/local/src/cmake.tar.gz
