// check excel
// check excel by rules define outside.
package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Usage: %s \n", os.Args[0])
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	usage()

	if len(os.Args) < 1 {
		usage()
	}

	fmt.Printf("End\n")
}
