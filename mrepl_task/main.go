//batch replace

package main

import (
	"flag"
	"github.com/mabetle/mcore"
	"strings"
	"os"
	"fmt"
)

// args
var (
	rootDir		string
	skipDir		string //not use, skip all dir start with .
	searchStr	string
	replaceStr	string
	verbose		bool
	exts		string
	recursive	bool
	help		bool
)

func usage(){
	fmt.Fprintf(os.Stderr,
		"Usage:%s [-d root dir] [-e exts] [-V verbose] [-r recursive] [-h]\"search content\" \"replace content\"\n\n",
		os.Args[0])

	flag.PrintDefaults()
}


func DoFlag(){
	flag.Usage=usage
	wd,_:=os.Getwd()
	flag.StringVar(&rootDir, "d",wd,
		"Set root dir, which dir to begin search and replace.")
	flag.StringVar(&exts, "e","go",
		"Extends, separate by comma for multiple file extends, dot can be ignored")
	flag.BoolVar(&recursive, "r", true,
		"Recursive dir or not, default is true")
	flag.BoolVar(&verbose, "V", true,
		"Print more info what is app doing.")
	flag.BoolVar(&help, "h", false,
		"Show help")

	flag.Parse()

}

func ShowArgs(){
	fmt.Println("App arguments")

	fmt.Println("Root dir           :", rootDir)
	fmt.Println("File extends       :", exts)
	fmt.Println("Verbose            :", verbose)
	fmt.Println("recursive          :", recursive)
	fmt.Println("")
	fmt.Println("Search content     :", searchStr)
	fmt.Println("Replace content    :", replaceStr)
}


func ScanText(msg string)string{
	return mcore.ReadNotBlankLineWithMsg(msg)
}

func ScanSearchContent(){
	searchStr = ScanText("Input search content:")
}

func ScanReplaceConent(){
	replaceStr = ScanText("Input replace content:")
}

func main() {
	DoFlag()

	if help{
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
		ScanReplaceConent()
	case 2:
		searchStr = flag.Args()[0]
		replaceStr = flag.Args()[1]
	default:
	}

	ShowArgs()
	replace()
}

func replace(){
	files:=mcore.GetSubFiles(rootDir, recursive, exts)
	fmt.Printf("Found %d files.\n", len(files))

	for _, item := range files {
		text,err :=mcore.ReadFileAll(item)
		if nil!=err{
			continue
		}

		if !strings.Contains(text, searchStr){
			if verbose{
				//fmt.Printf("File: %s not found matches\n" , item )
			}
			continue
		}else{
			nums:=strings.Count(text, searchStr)
			fmt.Printf("File: %s found %d matches.\n", item, nums)
			//do replace
			text = strings.Replace(text, searchStr, replaceStr, -1)
			mcore.WriteFile(item, text)
			fmt.Println("Write file :" , item)
		}

		//found
		data, err:=mcore.ReadFileLines(item)

		if err!=nil{
			fmt.Println(err)
			continue
		}

		for lineNum, line := range data {
			if strings.Contains(line, searchStr){
				fmt.Printf("%d %s\n",lineNum, line)
			}
		}
	}
}

