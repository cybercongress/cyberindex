FROM ubuntu:20.04

ENV GO_VERSION '1.18.10'
ENV GO_ARCH 'linux-amd64'
ENV GO_BIN_SHA '5e05400e4c79ef5394424c0eff5b9141cb782da25f64f79d54c98af0a37f8d49'

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

ENTRYPOINT ["./entrypoint.sh"]

CMD ["./start_script.sh"]
