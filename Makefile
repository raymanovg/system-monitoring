GOOS=darwin
GOARCH=amd64
CGO_ENABLED=0

BIN := "./bin/$(GOOS)/$(GOARCH)/monitoring"

VERSION?=$(shell git log --date=short --pretty=format:'%ad-%h' -n 1 | sed -E 's/([0-9]{4})\-([0-9]{2})\-([0-9]{2})/\1.\2.\3/g')
LDFLAGS := -X main.VERSION=$(VERSION)

build:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/monitoring

build-linux-img:
	docker build -t "linux/monitoring" -f build/Dockerfile .

build-linux-bin:
	GOOS=linux GOARCH=amd64 go build -v -o ./bin/linux/amd64/monitoring -ldflags "$(LDFLAGS)" ./cmd/monitoring

run-linux-monitoring: build-linux-bin build-linux-img
	docker run linux/monitoring:latest

.PHONY: build