FROM golang:1.24.0-alpine as base
WORKDIR /root

RUN apk add git curl make

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY pkg pkg
COPY templates templates
COPY web web
COPY Makefile Makefile
RUN mkdir static


# Build service in seperate stage.
FROM base as builder
RUN go build -o openslides-projector-service cmd/projectord/main.go

FROM node:22.15 as builder-web
COPY web web
COPY Makefile Makefile
RUN make build-web-assets


# Test build.
FROM base as testing

RUN apk add build-base

CMD go vet ./... && go test -test.short ./...


# Development build.
FROM base as development

COPY --from=builder-web /static ./static

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]
EXPOSE 9051

CMD CompileDaemon -log-prefix=false -build="go build -o openslides-projector-service cmd/projectord/main.go" -command="./openslides-projector-service"


# Productive build
FROM alpine:3

LABEL org.opencontainers.image.title="OpenSlides Projector Service"
LABEL org.opencontainers.image.description="The Projector Service is a http endpoint that serves projectors in Openslides."
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/OpenSlides/openslides-projector-service"

COPY --from=builder /root/openslides-projector-service .
COPY --from=builder-web /root/openslides-projector-service/static ./static
EXPOSE 9051
ENTRYPOINT ["/openslides-projector-service"]
