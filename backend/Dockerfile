FROM golang:1.16.14

WORKDIR /usr/src/app

COPY go.mod go.sum ./

ENV HOSTIP=192.168.1.11

RUN go mod download

COPY . .

RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]