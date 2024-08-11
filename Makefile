GO := GO111MODULE=on go

.PHONY: ci
ci:
	$(GO) mod tidy && \
	$(GO) mod download && \
	$(GO) mod verify && \
	$(GO) mod vendor && \
	$(GO) fmt ./... \

.PHONY: build
build:
	$(GO) build -mod=vendor -a -installsuffix cgo -tags musl -o main ./cmd/main.go

start:
	go run --tags dynamic $(shell pwd)/cmd/main.go

dev: 
	nodemon --exec go run --tags dynamic $(shell pwd)/cmd/main.go --signal SIGTERM

.PHONY: clean
clean:
	@rm -rf main ./vendor

test:
	go test -cover ./... -coverprofile=cover.out -covermode count && go tool cover -html cover.out -o cover.html

run:
	go run ./cmd/main.go
