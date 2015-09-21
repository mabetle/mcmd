package main

import (
	"mabetle/hub"
)

var (
	sql = hub.NewRootSql()
)

func main() {
	sql.CreateDatabase("demo")
}
