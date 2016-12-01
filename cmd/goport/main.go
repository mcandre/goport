package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"

	"github.com/docopt/docopt-go"
	"github.com/jhoonb/archivex"
	"github.com/mcandre/goport"
)

const Usage = `Usage:
  goport [options]
  goport -h | --help
  goport -v | --version

  Options:
    -a --application <name>  Specify an application name [default: $(pwd)]
    -l --label <name>        Specify a label, such as a version number
    -b --binaries <dir>      Specify a binary target directory [default: bin]
    -c --commands <dir>      Specify a command source directory [default: cmd]
    -h --help                Show usage information
    -v --version             Show version information`

const Perms = 0744

type BuildConfig struct {
	Banner string
	BinRoot string
	CmdRoot string
	Script string
	Target string
}

func build(buildConfig BuildConfig) {
	banner := buildConfig.Banner
	binRoot := buildConfig.BinRoot
	cmdRoot := buildConfig.CmdRoot
	script := buildConfig.Script
	target := buildConfig.Target

	osArchPattern := regexp.MustCompile("(.+)/(.+)")

	results := osArchPattern.FindStringSubmatch(target)

	if len(results) != 3 {
		panic(fmt.Sprintf("Error parsing target %s\n", target))
	}

	oSys, arch := results[1], results[2]

	suffix := ""

	if oSys == "windows" {
		suffix = ".exe"
	}

	executableBase := script + suffix

	branch := path.Join(binRoot, banner, oSys, arch)

	err := os.MkdirAll(branch, Perms)

	if err != nil {
		panic(err)
	}

	cmdDir := path.Join(cmdRoot, script)

	env := os.Environ()
	env = append(env, "GOOS="+oSys)
	env = append(env, "GOARCH="+arch)

	leaf := path.Join(branch, executableBase)

	log.Printf("Building %s\n", leaf)

	outputPath := path.Join("..", "..", leaf)

	command := exec.Command("go", "build", "-o", outputPath)
	command.Dir = cmdDir
	command.Env = env

	err = command.Run()

	if err != nil {
		panic(err)
	}
}

func main() {
	arguments, err := docopt.Parse(Usage, nil, true, goport.Version, false)

	if err != nil {
		panic(Usage)
	}

	app, _ := arguments["--application"].(string)

	if app == "$(pwd)" {
		cwd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		app = path.Base(cwd)
	}

	label, _ := arguments["--label"].(string)

	banner := app

	if label != "" {
		banner = fmt.Sprintf("%s-%s", app, label)
	}

	binRoot, _ := arguments["--binaries"].(string)

	cmdRoot, _ := arguments["--commands"].(string)

	targetBytes, err := exec.Command("go", "tool", "dist", "list").Output()

	if err != nil {
		panic(err)
	}

	targetLines := string(targetBytes)

	targetPattern := regexp.MustCompile("(?m)\\s+")

	targetsWithFinalEOL := targetPattern.Split(targetLines, -1)

	var targetsWithBetas []string

	for _, t := range targetsWithFinalEOL {
		if t != "" {
			targetsWithBetas = append(targetsWithBetas, t)
		}
	}

	targetBlacklist := regexp.MustCompile("(android/.+)|(darwin/arm64)")

	var targets []string

	for _, t := range targetsWithBetas {
		if !targetBlacklist.MatchString(t) {
			targets = append(targets, t)
		}
	}

	scriptEntriesWithBin, err := ioutil.ReadDir(cmdRoot)

	if err != nil {
		panic(err)
	}

	var scriptsWithJunkFiles []string

	binBase := path.Base(binRoot)

	for _, s := range scriptEntriesWithBin {
		name := s.Name()

		if name != binBase {
			scriptsWithJunkFiles = append(scriptsWithJunkFiles, name)
		}
	}

	var scripts []string

	junkFileBlacklist := regexp.MustCompile("Thumbs.db|.DS_Store")

	for _, s := range scriptsWithJunkFiles {
		if !junkFileBlacklist.MatchString(s) {
			scripts = append(scripts, s)
		}
	}

	for _, script := range scripts {
		for _, target := range targets {
			build(BuildConfig{
				Banner: banner,
				BinRoot: binRoot,
				CmdRoot: cmdRoot,
				Script: script,
				Target: target,
			})
		}
	}

	bannerDir := path.Join(binRoot, banner)
	archivePath := path.Join(binRoot, banner+".zip")

	log.Printf("Archiving ports to %s\n", archivePath)

	archive := new(archivex.ZipFile)
	defer archive.Close()

	archive.Create(archivePath)
	archive.AddAll(bannerDir, true)
}
