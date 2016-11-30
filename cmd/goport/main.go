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
    -a --application <name>  Specify an application name
    -l --label <name>        Specify a label, such as a version number
    -b --binaries <dir>      Specify a binary target directory
    -c --commands <dir>      Specify a command source directory
    -h --help                Show usage information
    -v --version             Show version information`

func main() {
	arguments, err := docopt.Parse(Usage, nil, true, goport.Version, false)

	if err != nil {
		panic(Usage)
	}

	application, applicationStatus := arguments["--application"].(string)

	if !applicationStatus {
		cwd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		application = path.Base(cwd)
	}

	log.Printf("application: %v\n", application)

	label, labelStatus := arguments["--label"].(string)

	log.Printf("label: %v\n", label)
	log.Printf("labelStatus: %v\n", labelStatus)

	if !labelStatus {
		label = ""
	}

	log.Printf("label: %v\n", label)
}
