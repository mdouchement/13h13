# build stage
FROM golang:alpine as build-env
LABEL maintainer="mdouchement"

RUN apk upgrade
RUN apk add --update --no-cache git curl

ENV CGO_ENABLED 0
ENV GO111MODULE on

WORKDIR /13h13
COPY . .

RUN go mod download
RUN go build -ldflags "-s -w" -o thirteen .

# final stage
FROM scratch
LABEL maintainer="mdouchement"

COPY --from=build-env /13h13/thirteen /usr/local/bin/

EXPOSE 8080
CMD ["thirteen"]
