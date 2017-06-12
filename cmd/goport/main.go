// Package main provides a goport executable.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"

	"github.com/jhoonb/archivex"
	"github.com/mcandre/goport"
)

var flagApplication = flag.String("application", "$(pwd)", "Application name")
var flagLabel = flag.String("label", "", "For example, a version number")
var flagBinaries = flag.String("binaries", "bin", "Binary output directory")
var flagCommands = flag.String("commands", "cmd", "Command source directory")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// Perms represents common Unix executable permissions.
const Perms = 0744

// BuildConfig presents customization points for gorport artifacts.
type BuildConfig struct {
	Banner  string
	BinRoot string
	CmdRoot string
	Script  string
	Target  string
}

// build generates cross-platform executables and archives them according to a given buildConfig.
func build(buildConfig BuildConfig) {
	banner := buildConfig.Banner
	binRoot := buildConfig.BinRoot
	cmdRoot := buildConfig.CmdRoot
	script := buildConfig.Script
	target := buildConfig.Target

	osArchPattern := regexp.MustCompile("(.+)/(.+)")

	results := osArchPattern.FindStringSubmatch(target)

	if len(results) != 3 {
		log.Panic("Error parsing target", target)
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
		log.Panic(err)
	}

	cmdDir := path.Join(cmdRoot, script)

	env := os.Environ()
	env = append(env, "GOOS="+oSys)
	env = append(env, "GOARCH="+arch)

	leaf := path.Join(branch, executableBase)

	log.Println("Building", leaf)

	outputPath := path.Join("..", "..", leaf)

	command := exec.Command("go", "build", "-o", outputPath)
	command.Dir = cmdDir
	command.Env = env

	err = command.Run()

	if err != nil {
		log.Panic(err)
	}
}

// main is the entrypoint for this application.
func main() {
	flag.Parse()

	switch {
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(1)
	case *flagVersion:
		fmt.Println(goport.Version)
		os.Exit(0)
	}

	if *flagApplication == "$(pwd)" {
		cwd, err := os.Getwd()

		if err != nil {
			log.Panic(err)
		}

		*flagApplication = path.Base(cwd)
	}

	banner := *flagApplication

	if *flagLabel != "" {
		banner = fmt.Sprintf("%s-%s", *flagApplication, *flagLabel)
	}

	targetBytes, err := exec.Command("go", "tool", "dist", "list").Output()

	if err != nil {
		log.Panic(err)
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

	scriptEntriesWithBin, err := ioutil.ReadDir(*flagCommands)

	if err != nil {
		log.Panic(err)
	}

	var scriptsWithJunkFiles []string

	binBase := path.Base(*flagBinaries)

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
				Banner:  banner,
				BinRoot: *flagBinaries,
				CmdRoot: *flagCommands,
				Script:  script,
				Target:  target,
			})
		}
	}

	bannerDir := path.Join(*flagBinaries, banner)
	archivePath := path.Join(*flagBinaries, banner+".zip")

	log.Print("Archiving ports to ", archivePath)

	archive := new(archivex.ZipFile)
	defer func() {
		if err := archive.Close(); err != nil {
			log.Panic(err)
		}
	}()

	if err := archive.Create(archivePath); err != nil {
		log.Panic(err)
	}
	if err := archive.AddAll(bannerDir, true); err != nil {
		log.Panic(err)
	}
}
