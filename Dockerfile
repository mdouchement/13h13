# build stage
FROM golang:alpine as build-env
MAINTAINER mdouchement

RUN apk upgrade

ENV CGO_ENABLED 0
ENV GO111MODULE on

WORKDIR /13h13
COPY . .

RUN go mod download
RUN go build -ldflags "-s -w" -o thirteen .

# final stage
FROM scratch
MAINTAINER mdouchement

COPY --from=build-env /13h13/thirteen /usr/local/bin/

EXPOSE 8080
CMD ["thirteen"]
