# Collect all functions that should be built and deployed.
functions := $(shell find cmd -name \*main.go | awk -F'/' '{print $$2}')

.PHONY: build clean deploy

upgrade:
	go get -u ./...

build:
	go mod tidy
	export GO111MODULE=on
	@for function in $(functions) ; do \
		mkdir -p bin/$$function ; \
		env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/$$function/bootstrap cmd/$$function/main.go ; \
		(cd bin/$$function && zip -FS bootstrap.zip bootstrap) ; \
	done

test:
	go test ./... -v

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	echo "See README.md for manual deploy, otherwise should go via CI/CD"
