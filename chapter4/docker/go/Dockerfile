FROM golang:1.22.4-bullseye

WORKDIR /go/src/app

ADD . /go/src/app

RUN go mod download

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz 