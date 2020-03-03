all: lint build test start

build:
	go build -o  _output/examples/libra-client  ./examples/libra-client

start: build
	_output/examples/libra-client

lint:
	find . -name '*.go' | xargs gofmt -w -s

test: 
	 go test ./libra

misspell:
	find . -name '*.go' -not -path './vendor/*' -not -path './_repos/*' | xargs misspell -error
