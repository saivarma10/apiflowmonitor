GO       := go
GOFMT    := $(GO) fmt
GOLINT   := $(GO) lint
GOCMD    := $(GO) run
GOFLAGS  := -mod=vendor

BINARY   := apimonitor
C_BINARY = transaction_monitor
transaction_monitor: cmd/c_program/transaction_monitor.c
	@gcc -o cmd/c_program/assets/$(C_BINARY) cmd/c_program/transaction_monitor.c -lcurl


.PHONY: all
all: build

.PHONY: build
build:
	$(GO) build -o $(BINARY) ./cmd

.PHONY: install
install:
	$(GO) install ./...

.PHONY: run
run:
	$(GOCMD) ./cmd/main.go

.PHONY: clean
clean:
	rm -f $(BINARY)

.PHONY: release
release: clean build
	tar -czf $(BINARY)-release.tar.gz $(BINARY)

.PHONY: vendor
vendor:
	$(GO) mod vendor
