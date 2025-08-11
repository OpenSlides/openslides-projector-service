FROM golang:1.24.3-alpine as base
WORKDIR /root/openslides-projector-service

RUN apk add git curl make

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY pkg pkg
COPY templates /root/templates
COPY web web
COPY locale locale
COPY Makefile Makefile
RUN mkdir static


# Build service in seperate stage.
FROM base as builder
RUN go build -o openslides-projector-service cmd/projectord/main.go

FROM node:24.5 as builder-web
COPY web web
COPY Makefile Makefile
RUN make build-web-assets


# Test build.
FROM base as testing

RUN apk add build-base

CMD go vet ./... && go test -test.short ./...


# Development build.
FROM base as development
WORKDIR /root

COPY --from=builder-web /static ./static

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]
EXPOSE 9051

CMD CompileDaemon -log-prefix=false -include="*.html" -build="go build -o projector-service ./openslides-projector-service/cmd/projectord/main.go" -command="./projector-service"


# Productive build
FROM alpine:3

LABEL org.opencontainers.image.title="OpenSlides Projector Service"
LABEL org.opencontainers.image.description="The Projector Service is a http endpoint that serves projectors in Openslides."
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/OpenSlides/openslides-projector-service"

COPY --from=builder /root/openslides-projector-service/openslides-projector-service .
COPY --from=builder-web /root/openslides-projector-service/static ./static
EXPOSE 9051
ENTRYPOINT ["/openslides-projector-service"]
