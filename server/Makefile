.PHONY: clean build format test

# -------------------------------------------------------------------------------- #
# Variables.

# Common.
OUTPUT_DIR = bin
SERVER_BIN = $(OUTPUT_DIR)/server

# Compiler and tools.
GO = go

# -------------------------------------------------------------------------------- #
# Goals.

all: build

build: format
	$(GO) build -o $(SERVER_BIN) ./app

clean:
	rm -fr ${OUTPUT_DIR}
	mkdir -p ${OUTPUT_DIR}

format:
	$(GO) mod tidy
	$(GO) fmt ./...

run: build
	./$(SERVER_BIN)

test:
	$(GO) test -v --cover ./...

# -------------------------------------------------------------------------------- #
