//batch replace

package main

import (
	"flag"
	"fmt"
	"github.com/mabetle/mcore"
	"os"
	"strings"
)

// args
var (
	rootDir    string
	skipDirs   string
	skipFiles  string
	searchStr  string
	replaceStr string
	exts       string
	verbose    bool
	recursive  bool
	help       bool
)

func usage() {
	fmt.Fprintf(os.Stderr,
		"Usage:%s [flags] \"search content\" \"replace content\"\n\n",
		os.Args[0])

	flag.PrintDefaults()
}

func DoFlag() {
	flag.Usage = usage
	wd, _ := os.Getwd()
	flag.StringVar(&rootDir, "d", wd,
		"Set root dir, which dir to begin search and replace.")
	flag.StringVar(&exts, "e", "",
		"Extends, separate by comma for multiple file extends, dot can be ignored")
	flag.StringVar(&skipDirs, "sd", "",
		"Skip Dirs, separate by comma for skip dirs")
	flag.StringVar(&skipFiles, "sf", "",
		"Skip Files, separate by comma for skip files")
	flag.BoolVar(&recursive, "r", true,
		"Recursive dir or not, default is true")
	flag.BoolVar(&verbose, "V", false,
		"Print more info what is app doing.")
	flag.BoolVar(&help, "h", false,
		"Show help")

	flag.Parse()

}

func ShowArgs() {
	fmt.Println("App Arguments")
	fmt.Println("       Root Dir :", rootDir)
	fmt.Println("   File Extends :", exts)
	fmt.Println("      Skip Dirs :", skipDirs)
	fmt.Println("     Skip Files :", skipFiles)
	fmt.Println("        Verbose :", verbose)
	fmt.Println("      Recursive :", recursive)
	fmt.Println(" Search Content :", searchStr)
	fmt.Println("Replace Content :", replaceStr)
}

func ScanSearchContent() {
	searchStr = mcore.ReadNotBlankLineWithMsg("Input Search Content:")
}

func ScanReplaceConent() {
	replaceStr = mcore.ReadLineWithMsg("Input Replace Content:")
}

func main() {
	DoFlag()

	if help {
		fmt.Println("Help about command")
		usage()
		return
	}

	switch flag.NArg() {
	case 0:
		ScanSearchContent()
		ScanReplaceConent()
	case 1:
		searchStr = flag.Args()[0]
		// replace str is ""
	case 2:
		searchStr = flag.Args()[0]
		replaceStr = flag.Args()[1]
	default:
	}

	ShowArgs()
	replace()
}

func replace() {
	files := mcore.GetSubFiles(rootDir, recursive, exts, skipDirs, skipFiles)
	fmt.Printf("Found %d files.\n", len(files))
	for _, item := range files {
		if verbose {
			show(item)
		}
		fileReplace(item, searchStr, replaceStr)
	}
}

func fileReplace(item, searchStr, replaceStr string) {
	text, err := mcore.ReadFileAll(item)
	// cannot read file
	if nil != err {
		fmt.Printf("Read file error: %v\n", err)
		return
	}
	if !strings.Contains(text, searchStr) {
		return
	}
	nums := strings.Count(text, searchStr)
	fmt.Printf("File: %s found %d matches.\n", item, nums)
	//do replace
	text = strings.Replace(text, searchStr, replaceStr, -1)
	if _, err := mcore.WriteFile(item, text); err != nil {
		fmt.Printf("Write file %s \n Error: %v", item, err)
	} else {
		fmt.Println("Write file :", item)
	}
}

func show(item string) {
	//found
	data, err := mcore.ReadFileLines(item)
	if err != nil {
		fmt.Println(err)
		return
	}
	for lineNum, line := range data {
		if strings.Contains(line, searchStr) {
			fmt.Printf("%d %s\n", lineNum, line)
		}
	}
}
