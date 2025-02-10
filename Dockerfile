FROM golang:1.23.6-alpine as base
WORKDIR /root

RUN apk add git curl make

ADD https://esbuild.github.io/dl/latest esbuild-install
RUN sh esbuild-install
RUN rm esbuild-install
RUN mv esbuild /usr/bin

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

RUN make build-web-assets


# Test build.
FROM base as testing

RUN apk add build-base

CMD go vet ./... && go test -test.short ./...


# Development build.
FROM base as development

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
EXPOSE 9051
ENTRYPOINT ["/openslides-projector-service"]
