FROM golang:latest

ARG JUNO_WORKERS=1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make build

RUN cp ./build/cyberindex /usr/local/bin/

CMD cyberindex parse config.toml 