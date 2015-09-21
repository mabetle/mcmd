package main

import (
	"github.com/mabetle/mcell/wxlsx"
)

func main() {
	location := "/Desktop/data/2015-SV.xlsx"
	wxlsx.PrintSheetByIndex(location, 0)
}
