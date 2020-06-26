GOBIN := go
BINNAME := httpd

all: test build

.PHONY: test
test:
	$(GOBIN) mod tidy
	$(GOBIN) test -v -race -cover ./...

.PHONY: build
build:
	$(GOBIN) build -o $(BINNAME) .

.PHONY: clean
clean:
	$(GOBIN) clean
	rm -f $(BINNAME)
