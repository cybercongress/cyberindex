FROM ubuntu:20.04

ENV GO_VERSION '1.16.5'
ENV GO_ARCH 'linux-amd64'
ENV GO_BIN_SHA 'b12c23023b68de22f74c0524f10b753e7b08b1504cb7e417eccebdd3fae49061'

WORKDIR /app

COPY go.mod go.sum ./

RUN apt-get update && apt-get install -y --no-install-recommends make build-essential wget ca-certificates

RUN wget -O go.tgz https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    echo "${GO_BIN_SHA} *go.tgz" | sha256sum -c - && \
    tar -C /usr/local -xzf go.tgz &&\
    rm go.tgz

ENV PATH="/usr/local/go/bin:$PATH"

COPY . .

RUN make build

RUN cp ./build/cyberindex /usr/local/bin/

CMD cyberindex parse config.toml 