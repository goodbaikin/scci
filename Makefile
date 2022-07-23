default: install

.PHONY: build
build:
	CGO_EMABLED=0 go build -o ./cmd/scmav/ -ldflags '-s -w' ./cmd/scmav

.PHONY: install
install:
	go install ./cmd/scmav
