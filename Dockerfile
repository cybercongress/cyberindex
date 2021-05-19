FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download \ 
&& apt install make

COPY . .

RUN make build

RUN cp ./build/cyberindex /usr/local/bin/

CMD cyberindex parse config.toml 