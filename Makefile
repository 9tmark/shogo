.PHONY: clean build
define ReadTxt
$(shell cat $(1))
endef
DOMAIN:=$(call ReadTxt,domain.txt)
KEY:=$(call ReadTxt,key.txt)
VERSION:=$(call ReadTxt,VERSION)

clean:
	rm -rf ./build/*; \
	rm ./*.out; \
	go clean

build:
	mkdir -p ./build && \
	go build -ldflags "-X main.Domain=${DOMAIN}" -ldflags "-X main.Key=${KEY}" -ldflags "-X main.Version=${VERSION}" -o=./build ./cmd/shogo.go

install: clean build
	chmod +x ./build/shogo && \
	sudo cp ./build/shogo /usr/local/bin/ && \
	sudo ln -s /usr/local/bin/shogo /usr/local/bin/shg

uninstall:
	sudo rm -rf /usr/local/bin/shogo && \
	sudo rm -rf /usr/local/bin/shg
