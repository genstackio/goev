BINARY_NAME	= goev

all: install

build:
	@go build -o $(BINARY_NAME)

install:
	@go mod tidy