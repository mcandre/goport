# goport - build bots ♥ me

# EXAMPLE

```
$ goport -l 0.0.1
...
Archiving ports to bin/goport-0.0.1.zip

$ unzip -l bin/goport-0.0.1.zip
Archive:  bin/goport-0.0.1.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
        0  11-30-2016 15:38   goport-0.0.1/
        0  11-30-2016 14:50   goport-0.0.1/darwin/
        0  11-30-2016 15:45   goport-0.0.1/darwin/amd64/
  2648144  11-30-2016 15:45   goport-0.0.1/darwin/amd64/goport
...
  2795520  11-30-2016 15:51   goport-0.0.1/windows/amd64/goport.exe
---------                     -------
 79477875                     73 files
```

See [mcandre/go-ios7crypt](https://github.com/mcandre/go-ios7crypt/blob/master/Makefile#L32) for more examples.

# ABOUT

`goport` is a light wrapper around `go build`, an amazing cross-compiler! While developers can manually port applications with individual `env GOOS=... GOARCH=... go build ...` commands, `goport` abstracts certain platform-dependent details:

* Target all supported platforms, from `windows/386` to `darwin/amd64`
* Support multiple `cmd/*` scripts
* Accept host-specific source paths like `~` and `C:\\`
* Respect platform-specific binary naming conventions like `*.exe`
* Organize ports in a simple directory tree (`<banner>/<os>/<arch>/<app>[.suffix]>`)
* Package binaries in a single `.zip` file for easy HTTP hosting

# DOWNLOAD

https://github.com/mcandre/goport/releases

# REQUIREMENTS

* [Go](https://golang.org) 1.7+ with [$GOPATH configured](https://gist.github.com/mcandre/ef73fb77a825bd153b7836ddbd9a6ddc)
* [zip](https://linux.die.net/man/1/zip)

## Optional

* [Git](https://git-scm.com)
* [Make](https://www.gnu.org/software/make/)
* [Bash](https://www.gnu.org/software/bash/)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/goport/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone git@github.com:mcandre/goport.git $GOPATH/src/github.com/mcandre/goport
$ cd $GOPATH/src/github.com/mcandre/goport
$ git submodule update --init --recursive
$ sh -c 'cd cmd/goport && go install'
```

# TEST

```
$ make integration-test
```

# PORT

```
$ make port
```

# LINT

Keep the code tidy:

```
$ make lint
```

# GIT HOOKS

See `hooks/`.
