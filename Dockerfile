FROM golang:alpine


RUN export GOPATH=usr/local/go

RUN mkdir /go/src/app-login
WORKDIR /go/src/app-login

COPY . .

RUN go env

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates

RUN go get -d -v github.com/go-ldap/ldap/v3

EXPOSE 8000