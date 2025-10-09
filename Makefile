build-dev:
	docker build . --target development --tag openslides-projector-dev

run-tests:
	docker build . --target testing --tag openslides-projector-test
	docker run openslides-projector-test

all: gofmt gotest golinter

gotest:
	go test ./...

golinter:
	golint -set_exit_status ./...

gofmt:
	gofmt -l -s -w .

gogenertate:
	go generate ./...

nodelinter: | install-web-asset-deps
	cd web && npm run lint

install-web-asset-deps:
	cd web && npm i

build-live-all:
	make build-watch-web-assets &
	make build-live

build-live:
	go run github.com/githubnemo/CompileDaemon@v1.4.0 -log-prefix=false -include="*.html" -build="go build -o projector-service ./cmd/projectord/main.go" -command="./projector-service"

build-web-assets: | install-web-asset-deps
	cd web && npm run build

build-watch-web-assets: | install-web-asset-deps
	cd web && npm run build-watch
