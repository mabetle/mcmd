// gofind cmd
// Usage: gofind [-d -e -r -c -sd -sf] searchContent
package main

import (
	"flag"
	"fmt"
	"github.com/mabetle/mcore"
	"github.com/mabetle/mcore/mcon"
	"os"
	"strings"
)

var (
	dir           string
	exts          string
	skipDirs      string
	skipFiles     string
	searchContent string
	verbose       bool
	recursive     bool
	help          bool
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [Flags] Search Content\n", os.Args[0])
	flag.PrintDefaults()
}

func DoFlag() {
	wd, _ := os.Getwd()
	flag.StringVar(&dir, "d", wd,
		"Search dir")
	flag.StringVar(&exts, "e", "",
		"File extends")
	flag.StringVar(&skipDirs, "sd", "",
		"Skip Dirs, separate by comma for skip dirs")
	flag.StringVar(&skipFiles, "sf", "",
		"Skip Files, separate by comma for skip files")
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
	fmt.Printf("      Root Dir : %s\n", dir)
	fmt.Printf("  File Extends : %s\n", exts)
	fmt.Printf("     Skip Dirs : %s\n", skipDirs)
	fmt.Printf("    Skip Files : %s\n", skipFiles)
	fmt.Printf("     Recursive : %v\n", recursive)
	fmt.Printf("Search Content : %s\n", searchContent)
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
	if flag.NArg() > 0 {
		searchContent = strings.Join(flag.Args(), " ")
	}

	// check searchContent
	if searchContent == "" {
		searchContent = mcore.ReadNotBlankLineWithMsg("Input Search Content:")
	}

	ShowArgs()
	Search(dir, exts, recursive, skipDirs, skipFiles, searchContent)
}

func Search(path string, exts string, recursive bool, skipDirs, skipFiles, content string) {
	files := mcore.GetSubFiles(path, recursive, exts, skipDirs, skipFiles)
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
				fmt.Printf("%d ", lineNum+1)
				lineA := mcore.String(line).Split(content)
				for i, v := range lineA {
					fmt.Printf(v)
					if i != len(lineA)-1 {
						mcon.PrintGreen(content)
					}
				}
				fmt.Println()
			}
		}
	}
}
