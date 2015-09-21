package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func usage() {
	fmt.Printf("Usage: prefix_task [add | remove] prefix [ext] \n")
}

type FileNameWalkFunc func(fn string)

func DirFilesWalk(dir string, ext string, walkFunc func(fn string)) {
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
		return
	}
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		if ext != "" && !strings.HasSuffix(f.Name(), ext) {
			continue
		}
		walkFunc(f.Name())
	}
}

func addPrefix(pwd, prefix, ext string) {
	walkFunc := func(fn string) {
		newName := fmt.Sprintf("%s_%s", prefix, fn)
		errMove := os.Rename(fn, newName)
		if errMove != nil {
			log.Println(errMove)
		}
	}
	DirFilesWalk(pwd, ext, walkFunc)
}

func removePrefix(pwd, prefix, ext string) {
	walkFunc := func(fn string) {
		newName := strings.TrimPrefix(fn, prefix+"-")
		newName = strings.TrimPrefix(newName, prefix+"_")
		newName = strings.TrimPrefix(newName, prefix+".")
		errMove := os.Rename(fn, newName)
		if errMove != nil {
			log.Println(errMove)
		}
	}
	DirFilesWalk(pwd, ext, walkFunc)
}

func lowerExt(pwd, ext string) {
	walkFunc := func(fn string) {
		fileNA := strings.Split(fn, ".")
		// no ext
		if len(fileNA) < 2 {
			return
		}

		ext := fileNA[len(fileNA)-1]
		extLower := strings.ToLower(ext)

		newName := strings.TrimSuffix(fn, ext) + extLower

		errMove := os.Rename(fn, newName)
		if errMove != nil {
			log.Println(errMove)
		}
	}
	DirFilesWalk(pwd, ext, walkFunc)
}

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		usage()
		return
	}
	pwd, errPwd := os.Getwd()
	if errPwd != nil {
		log.Println(errPwd)
		return
	}

	action := flag.Arg(0)

	// lower all ext
	if action == "lowerExt" {
		lowerExt(pwd, "")
		return
	}

	prefix := flag.Arg(1)
	ext := ""
	if flag.NArg() == 3 {
		ext = flag.Arg(2)
	}

	switch action {
	case "add":

		addPrefix(pwd, prefix, ext)
	case "remove", "delete":

		removePrefix(pwd, prefix, ext)
	case "update", "change":

		fmt.Printf("TODO\n")
	default:
		fmt.Printf("Unknow action: %s\n", action)
		usage()
	}
}
