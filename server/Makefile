.PHONY: clean build format test

OUTPUT_DIR = bin


all: clean build

clean:
	rm -fr ${OUTPUT_DIR}/*

build: format
	go build -o ${OUTPUT_DIR}/server ./app

format:
	go mod tidy
	go fmt ./...

test:
	go test -v --cover ./...
