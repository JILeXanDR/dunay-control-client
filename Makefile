OUT := "bin/client_$$(git rev-parse --short HEAD)"

default: build

build:
	rm -R bin/client_*
	go build -o $(OUT)

run:
	go run .