VERSION=0.0.1

.PHONY: port clean clean-ports

all: integration-self-test

integration-self-test: port

port:
	goport -a goport -l $(VERSION) -b bin -c cmd

govet:
	go list ./... | grep -v vendor | xargs go vet -v

gofmt:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;

goimport:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec goimports -w {} \;

lint: govet gofmt goimport

clean: clean-ports

clean-ports:
	rm -rf bin
