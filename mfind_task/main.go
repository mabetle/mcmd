// gofind cmd
// Usage: gofind --dir --ext searchContent
package main

import (
	"flag"
	"fmt"
	"github.com/mabetle/mcore"
	"os"
	"strings"
	"github.com/mabetle/mcore/mcon"
)

var (
	dir           string
	exts          string
	searchContent string
	verbose       bool
	recursive     bool
	help          bool
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [-d dir] [-e extends] [-V] [-r] content \n", os.Args[0])
	flag.PrintDefaults()
}

func DoFlag() {
	wd, _ := os.Getwd()
	flag.StringVar(&dir, "d", wd,
		"Search dir")
	flag.StringVar(&exts, "e", "go",
		"File extends")
	flag.BoolVar(&verbose, "V", false,
		"Verbose")
	flag.BoolVar(&recursive, "r", true,
		"Recursive")
	flag.BoolVar(&help, "h", false,
		"Show help")

	flag.Usage = usage
	flag.Parse()
}

func ShowArgs() {
	fmt.Println()
	//show vars
	fmt.Printf("Root dir      :%s\n", dir)
	fmt.Printf("File extend   :%s\n", exts)
	fmt.Printf("Search Content:%s\n", searchContent)
	fmt.Println()
}

func main() {
	DoFlag()

	if help {
		fmt.Println("Help about command")
		usage()
		return
	}

	// should tell me what to search
	if flag.NArg() == 1 {
		searchContent = flag.Args()[0]
	}

	// check searchContent
	if searchContent == "" {
		searchContent = mcore.ReadNotBlankLineWithMsg("Input Search Content:")
	}

	ShowArgs()
	Search(dir, exts, searchContent)
}

func Search(path string, exts string, content string) {
	files := mcore.GetSubFiles(path, true, exts)
	for _, item := range files {
		text, err := mcore.ReadFileAll(item)
		if nil != err {
			continue
		}

		if !strings.Contains(text, content) {
			if verbose {
				fmt.Printf("File: %s not found matches\n", item)
			}
			continue
		} else {
			nums := strings.Count(text, content)
			fmt.Printf("File: %s found %d matches.\n", item, nums)
		}

		//found
		data, err := mcore.ReadFileLines(item)

		if err != nil {
			fmt.Println(err)
			continue
		}

		for lineNum, line := range data {
			if strings.Contains(line, content) {
				fmt.Printf("%d ", lineNum)
				lineA:=mcore.String(line).Split(content)
				for i, v := range  lineA{
					fmt.Printf(v)
					if i !=len(lineA)-1{
						mcon.PrintGreen(content)
					}
				}
				fmt.Println()
			}
		}
	}
}
