package main

import (
	"log"
	"os"
	"path"

	"github.com/docopt/docopt-go"
	"github.com/mcandre/goport"
)

const Usage = `Usage:
  goport [options]
  goport -h | --help
  goport -v | --version

  Options:
    -a --application <name>  Specify an application name [default: $(pwd)]
		-l --label <name>        Specify a label, such as a version number [default: nil]
		-b --binaries <dir>      Specify a binary target directory [default: bin]
		-c --commands <dir>      Specify a command source directory [default: cmd]
    -h --help                Show usage information
    -v --version             Show version information`

func main() {
	arguments, err := docopt.Parse(Usage, nil, true, goport.Version, false)

	if err != nil {
		panic(Usage)
	}

	application, _ := arguments["--application"].(string)

	if application == "$(pwd)" {
		cwd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		application = path.Base(cwd)
	}

	log.Printf("application: %v\n", application)

	label, _ := arguments["--label"].(string)

	log.Printf("label: %v\n", label)

	binariesDirectory, _ := arguments["--binaries"].(string)

	log.Printf("binariesDirectory: %v\n", binariesDirectory)
}
