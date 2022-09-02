default: install

.PHONY: build
build:
	CGO_EMABLED=0 go build -o ./cmd/scci/ -ldflags '-s -w' ./cmd/scci

.PHONY: install
install: build
	go install ./cmd/scci
