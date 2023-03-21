# Makefile for outsert

BINARY_NAME=outsert

test:
	go test -v .

build:
	go build -o $(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

.PHONY: test build clean
