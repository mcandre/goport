VERSION=0.0.1

.PHONY: port clean clean-ports

all: integration-self-test

integration-self-test: port

port:
	goport -l $(VERSION)

govet:
	go list ./... | grep -v vendor | xargs go vet -v

golint:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec golint {} \;

gofmt:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;

goimport:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec goimports -w {} \;

bashate:
	find . \( -wholename '*/.git/*' -o -wholename '*/node_modules*' -o -name '*.bat' \) -prune -o -type f \( -wholename '*/lib/*' -o -wholename '*/hooks/*' -o -name '*.sh' -o -name '*.bashrc*' -o -name '.*profile*' -o -name '*.envrc*' \) -print | xargs bashate

shlint:
	find . \( -wholename '*/.git/*' -o -wholename '*/node_modules*' -o -name '*.bat' \) -prune -o -type f \( -wholename '*/lib/*' -o -wholename '*/hooks/*' -o -name '*.sh' -o -name '*.bashrc*' -o -name '.*profile*' -o -name '*.envrc*' \) -print | xargs shlint

checkbashisms:
	find . \( -wholename '*/.git/*' -o -wholename '*/node_modules*' -o -name '*.bat' \) -prune -o -type f \( -wholename '*/lib/*' -o -wholename '*/hooks/*' -o -name '*.sh' -o -name '*.bashrc*' -o -name '.*profile*' -o -name '*.envrc*' \) -print | xargs checkbashisms -n -p

shellcheck:
	find . \( -wholename '*/.git/*' -o -wholename '*/node_modules*' -o -name '*.bat' \) -prune -o -type f \( -wholename '*/lib/*' -o -wholename '*/hooks/*' -o -name '*.sh' -o -name '*.bashrc*' -o -name '.*profile*' -o -name '*.envrc*' \) -print | xargs shellcheck

editorconfig:
	sh editorconfig.sh

lint: govet golint gofmt goimport bashate shlint checkbashisms shellcheck editorconfig

clean: clean-ports

clean-ports:
	rm -rf bin
