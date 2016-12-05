VERSION=0.0.1

.PHONY: port clean clean-ports

all: integration-self-test

integration-self-test: port

port:
	goport -l $(VERSION)

govet:
	go list ./... | grep -v vendor | xargs go vet -v

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
	find . \( -wholename '*/bin/*' -o -wholename '*/ts/*.js' -o -wholename '*/dash/lib/*' -o -wholename '*/ash/lib/*' -o -name '*.bc' -o -name '*.aux' -o -name '*.jad' -o -name '*.m' -o -name '*.snu' -o -name '*.txt' -o -name '*.md' -o -name '*.rkt' -o -name '*.clj' -o -name '*.lsp' -o -name .yaws -o -name '*.pdf' -o -name '*.ps' -o -wholename '*/.idea/*' -o -name '*.iml' -o -name '*.ser' -o -name '*.[ps]k' -o -name '*.flip' -o -name '*.db' -o -name '*.log' -o -wholename '*/bower_components/*' -o -wholename '*/vendor/*' -o -wholename '*/*.xcodeproj/*' -o -wholename '*/*.dSYM/*' -o -wholename '*/build/*' -o -wholename '*/*.app/*' -o -name '*.scpt' -o -wholename '*/perl/Makefile' -o -wholename '*/CMakeFiles/*' -o -name '*.cmake' -o -name '*.lock' -o -name '*.cm[io]' -o -name '*.hi' -o -name '*.swiftdoc' -o -name '*.swiftmodule' -o -name '*.rlib' -o -name '*.dylib' -o -name '*.so' -o -name '*.o' -o -name '*.beam' -o -name '*.dump' -o -name '*.pyc' -o -name '*.jar' -o -name '*.class' -o -name '*.bin' -o -wholename '*/tmp/*' -o -name .gitmodules -o -wholename '*/.git/*' -o -wholename '*/node_modules/*' -o -wholename '*/.cabal/*' -o -name '*.ttf' -o -name '*.plist' -o -name '*.dot' -o -name '*.wav' -o -name '*.jpeg' -o -name '*.jpg' -o -name '*.ico' -o -name '*.png' -o -name '*.gif' -o -name .DS_Store -o -name Thumbs.db \) -prune -o -type f -print | pargs -n 100 node_modules/.bin/editorconfig-tools check

lint: govet gofmt goimport bashate shlint checkbashisms shellcheck editorconfig

clean: clean-ports

clean-ports:
	rm -rf bin
