FROM golang:1.15.12

WORKDIR /app

COPY go.mod go.sum ./

RUN apt install make

COPY . .

RUN make build

RUN cp ./build/cyberindex /usr/local/bin/

CMD cyberindex parse config.toml 