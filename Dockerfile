# build stage
FROM golang:1.11-alpine as build-env
MAINTAINER mdouchement

RUN apk upgrade
RUN apk add --update --no-cache git curl

RUN mkdir -p /go/src/github.com/mdouchement/13h13
WORKDIR /go/src/github.com/mdouchement/13h13

COPY . /go/src/github.com/mdouchement/13h13/

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure -v
RUN go build -o thirteen thirteen.go

# final stage
FROM alpine:3.8
MAINTAINER mdouchement

COPY --from=build-env /go/src/github.com/mdouchement/13h13/thirteen /usr/local/bin/

EXPOSE 8080
CMD ["thirteen"]
