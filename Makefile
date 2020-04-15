.DEFAULT_GOAL: NaviBot
.PHONY: rpi prep verify update run test rm

NaviBot: prep
	go build -o build/NaviBot main.go

rpi:
	GOOS="linux" GOARCH="arm" GOARM="6" go build -o build/NaviBot-RPi main.go

prep: verify
	mkdir -p build

verify:
	go mod verify

update:
	go get -u

run:
	go run main.go

test: NaviBot
	./build/NaviBot

rm:
	rm -r build/*
