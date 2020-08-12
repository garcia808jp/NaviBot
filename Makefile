.DEFAULT_GOAL: build
.PHONY: rpi cirno prep verify update run test rm

build: prep
	go build -o bin/NaviBot *.go

rpi: prep
	GOOS="linux" GOARCH="arm" GOARM="6" go build -o bin/NaviBot-RPi *.go

cirno: prep
	GOOS="plan9" GOARCH="amd64" go build -o bin/NaviBot-9f *.go

prep: verify
	mkdir -p bin

verify:
	go mod verify

update:
	go get -u

run:
	go run *.go

test: build
	./bin/NaviBot

rm:
	rm -r bin/*
