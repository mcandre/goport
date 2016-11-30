package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/mcandre/goport"
)

const Usage = `Usage:
  goport [-alvh]

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
		fmt.Println(Usage)
	}

	fmt.Printf("Arguments: %v\n", arguments)

	// ...
}
