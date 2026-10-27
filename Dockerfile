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

# Prepare Scratch for runtime and security purposes.
RUN echo "nobody:x:65534:65534:nobody:/_nonexistent:/bin/false" > /etc_passwd

# final stage
FROM scratch
LABEL maintainer="mdouchement"

COPY --from=build-env /13h13/thirteen /usr/local/bin/

# Setup Scratch image to be secure.
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /etc_passwd /etc/passwd
USER nobody

EXPOSE 8080
CMD ["thirteen"]
