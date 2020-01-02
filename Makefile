OUT := "bin/client_$$(git describe)"

default: build

build:
	go build -o $(OUT)

run:
	go run .