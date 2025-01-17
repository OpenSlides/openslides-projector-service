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

build-web-assets:
	esbuild web/projector.js web/projector.css web/slide/*.css web/slide/*.js --outdir=static/ --external:*.woff --bundle --minify --sourcemap --target=chrome58,firefox57,safari11,edge16

build-watch-web-assets:
	esbuild web/projector.js web/projector.css web/slide/*.css web/slide/*.js --outdir=static/ --external:*.woff --watch --bundle --minify --sourcemap --target=chrome58,firefox57,safari11,edge16
