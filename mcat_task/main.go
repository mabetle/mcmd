package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// cat file, args 1 as filename
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: mcat_task file...\n")
		return
	}
	Cat(os.Args[1:])
}

func Cat(args []string) {
	for _, v := range args {
		filename := v
		ba, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("%s\n", string(ba))
	}
}
