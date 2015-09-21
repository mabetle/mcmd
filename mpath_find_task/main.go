package main

import (
	"flag"
	"fmt"
	"github.com/mabetle/mcore"
	"os"
	"strings"
)

var (
	find = ""
)

func usage() {
	fmt.Println("Usage: mpath_find file")
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
		return
	}

	find = flag.Args()[0]

	path := os.Getenv("PATH")

	rs := GetPathFiles(path, find)

	// output dir
	for _, line := range rs {
		fmt.Printf("%s\n", line)
	}
}

func GetPathFiles(path string, pattern string) (fs []string) {
	path = strings.Replace(path, ";", ":", -1)
	sep := ":"
	pathA := strings.Split(path, sep)
	for _, path := range pathA {
		pfs := mcore.GetDirSubFiles(path)
		for _, pf := range pfs {
			if mcore.NewString(pf).IsHasSuffix("/" + pattern) {
				fs = append(fs, pf)
			}
		}
	}
	return
}
