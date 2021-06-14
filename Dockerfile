FROM docker.io/golang:stretch as build

ENV GO111MODULE on
ENV GOPROXY http://proxy.golang.org

RUN mkdir -p /src/token-bot
WORKDIR /src/token-bot

# Copy the Go Modules manifests
COPY go.mod .
COPY go.sum .

RUN go mod download

ADD . .

RUN DATE=$(date -u '+%Y-%m-%d_%I:%M:%S%p') \
    GOOS=linux GOARCH=amd64 \
        /usr/local/go/bin/go build ${FLAGS} \
        -tags release \
        -ldflags "-X ${PACKAGE}/cmd.Version=${VERSION} -X ${PACKAGE}/cmd.BuildDate=${DATE}" \
        -o /dist/token main.go

FROM registry.hub.docker.com/library/ubuntu:focal-20210416@sha256:86ac87f73641c920fb42cc9612d4fb57b5626b56ea2a19b894d0673fd5b4f2e9

RUN apt-get update && apt-get install -y \
    curl \
 && rm -rf /var/lib/apt/lists/*

COPY --from=build /dist/token /token
COPY config/config.yml /config/config.yml

ENTRYPOINT [ "/token" ]
