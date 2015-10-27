package main

import (
	"mabetle/libs/hubs"
	"fmt"
	"os"
	"flag"
)

var (
	sql = hubs.GetRootSql()
	dbname string
)

func doFlag(){
	flag.Usage = func(){
		fmt.Fprintf(os.Stderr, "Usage: %s dbname", os.Args[0])
	}
	flag.Parse()
}

func main() {
	doFlag()
	
	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	dbname = flag.Arg(0)

	sql.CreateDatabase(dbname)
}
