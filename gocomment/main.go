package main

import (
	"flag"
	"fmt"
	"os"
)

// args
var (
	file string
)

func doFlag() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s file", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}

func fixComment(file string) {
	// TODO
}

func main() {
	doFlag()
	if flag.NArg() < 1 {
		flag.Usage()
		return
	}
	
	file = flag.Arg(1)

	fixComment(file)
}
