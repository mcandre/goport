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
Optional

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
