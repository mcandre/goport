# goport - build bots ♥ me

# EXAMPLE

```
$ goport -l 0.0.1

$ unzip -l bin/goport-0.0.1.zip
  2648144  11-30-2016 15:45   goport-0.0.1/darwin/amd64/goport
  2284208  11-30-2016 15:45   goport-0.0.1/darwin/386/goport
  2315114  11-30-2016 15:45   goport-0.0.1/linux/386/goport
  2699164  11-30-2016 15:45   goport-0.0.1/linux/amd64/goport
  2376185  11-30-2016 15:45   goport-0.0.1/linux/arm/goport
  2934726  11-30-2016 15:45   goport-0.0.1/linux/arm64/goport
  3200543  11-30-2016 15:45   goport-0.0.1/linux/mips64/goport
  3200679  11-30-2016 15:46   goport-0.0.1/linux/mips64le/goport
  3002876  11-30-2016 15:46   goport-0.0.1/linux/ppc64/goport
  3002851  11-30-2016 15:46   goport-0.0.1/linux/ppc64le/goport
  3003563  11-30-2016 15:46   goport-0.0.1/linux/s390x/goport
  2420736  11-30-2016 15:51   goport-0.0.1/windows/386/goport.exe
  2795520  11-30-2016 15:51   goport-0.0.1/windows/amd64/goport.exe
...
```

More examples:

* [go-ios7crypt](https://github.com/mcandre/go-ios7crypt/tree/117ec78f571cbe3bb42313f6a9dd3f77a4aaa322)
* [go-chop](https://github.com/mcandre/go-chop/tree/579da02592cd3d95fd1b56692c9187fb919bac79)
* [go-hextime](https://github.com/mcandre/go-hextime/tree/48df8ab973694e15c2c87ee7b5e521af4f9174bf)
* [go-rotate](https://github.com/mcandre/go-rotate/tree/3bbf4d965631d3fd9606bdf9363e9c6476ac1423)
* [go-swatch](https://github.com/mcandre/go-swatch/commit/1f044fef9fb375e1b20a9d414289e686e70948f8)
* [karp](https://github.com/mcandre/karp/tree/e3713a5ed06f20d78f94e0362f391b0453e13241)

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

## Optional

* [coreutils](https://www.gnu.org/software/coreutils/coreutils.html)
* [make](https://www.gnu.org/software/make/)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [shlint](https://rubygems.org/gems/shlint)
* [shellcheck](http://hackage.haskell.org/package/ShellCheck)

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

# CREDITS

Shout out to [jhoonb/archivex](https://github.com/jhoonb/archivex) for simplifying recursive archiving!

# ALTERNATIVES

* [gox](https://github.com/mitchellh/gox) + [zipc](https://github.com/mcandre/zipc)
* [goport.sh](https://gist.github.com/mcandre/287a09b12f20d2781aa0875cb97c79fb)
