ARG CONTEXT=prod

FROM golang:1.25.4-alpine AS base

## Setup
ARG CONTEXT
WORKDIR /root/openslides-projector-service
ENV APP_CONTEXT=${CONTEXT}

## Install
RUN apk add --no-cache \
    curl \
    git \
    make

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY pkg pkg
COPY templates templates
COPY web web
COPY locale locale
COPY Makefile Makefile
RUN mkdir static


# Build service in seperate stage.
FROM base AS builder

RUN go build -o openslides-projector-service cmd/projectord/main.go


FROM node:22.13 AS builder-web

COPY web /web
COPY Makefile /Makefile
RUN make build-web-assets


# Test build.
FROM base AS tests

RUN apk add --no-cache \
    build-base

CMD go vet ./... && go test -test.short ./...


# Development build.
FROM base AS dev
WORKDIR /root/openslides-projector-service

RUN apk add --no-cache \
    nodejs \
    npm

COPY --from=builder-web /static ./static
COPY web web
RUN cd web && npm ci
COPY Makefile Makefile

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@v1.4.0"]
EXPOSE 9051

CMD ["make", "build-live-all"]


# Productive build
FROM alpine:3 AS prod

## Setup
ARG CONTEXT
ENV APP_CONTEXT=prod

LABEL org.opencontainers.image.title="OpenSlides Projector Service"
LABEL org.opencontainers.image.description="The Projector Service is a http endpoint that serves projectors in Openslides."
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/OpenSlides/openslides-projector-service"

COPY --from=builder /root/openslides-projector-service/openslides-projector-service /
COPY --from=builder /root/openslides-projector-service/templates /templates
COPY --from=builder-web /static /static
EXPOSE 9051
ENTRYPOINT ["/openslides-projector-service"]
