package main

import (
	"flag"
	"fmt"
	"mabetle/libs/hubs"
	"github.com/mabetle/mcore"
	"os"
)

func MakeDir(dir string) {
	mcore.MakeDir(dir)
}

func MakeOrgDir(path string) {
	for k, v := range hub.OrgNames {
		dir := fmt.Sprintf("%s/%d-%s", path, 100+k, v)
		MakeDir(dir)
	}
}

//
var wd string
var subCmd string

func Usage() {
	fmt.Printf("Usage: %s [mkorg|help] \n", os.Args[0])
	os.Exit(2)
}

func main() {
	flag.Usage = Usage
	flag.Parse()

	if flag.NArg() < 1 {
		Usage()
	}

	var err error
	wd, err = os.Getwd()
	if err != nil {
		fmt.Println("Get work dir error.")
	}

	// cmd args should more than one

	subCmd = os.Args[1]
	switch subCmd {
	case "mkorg":
		if wd != "" {
			MakeOrgDir(wd)
		}
	case "help":
		Usage()
	default:
		Usage()
	}
}
