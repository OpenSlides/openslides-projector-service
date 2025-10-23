override SERVICE=projector

# Build images for different contexts

build-prod:
	docker build ./ $(ARGS) --tag "openslides-$(SERVICE)" --build-arg CONTEXT="prod" --target "prod"

build-dev:
	docker build ./ $(ARGS) --tag "openslides-$(SERVICE)-dev" --build-arg CONTEXT="dev" --target "dev"

build-tests:
	docker build ./ $(ARGS) --tag "openslides-$(SERVICE)-tests" --build-arg CONTEXT="tests" --target "tests"

build-live-all:
	make build-watch-web-assets &
	make build-live

build-live:
	go run github.com/githubnemo/CompileDaemon@v1.4.0 -log-prefix=false -include="*.html" -build="go build -o projector-service ./cmd/projectord/main.go" -command="./projector-service"

install-web-asset-deps:
	cd web && npm i

build-web-assets: | install-web-asset-deps
	cd web && npm run build

build-watch-web-assets: | install-web-asset-deps
	cd web && npm run build-watch

# Tests
run-tests: | build-tests
	docker run openslides-projector-tests

lint: gofmt gotest golinter

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
