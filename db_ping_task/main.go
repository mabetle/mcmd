package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mabetle/mlog"
)

var (
	logger     = mlog.GetLogger("main")
	driverName = "mysql"
)

func PingDB(connURL string) {
	db, errDB := sql.Open(driverName, connURL)
	if errDB != nil {
		logger.Errorf("Open sql error.Error: %v", errDB)
		return
	}

	// infact if db not work, app still can go on.
	if err := db.Ping(); err != nil {
		logger.Errorf("Ping db error.Error: %v", err)
		return
	}
	logger.Info("Pingdb Success.")
}

func main() {
	// connURL demo: root:xxx@tcp(db.mabetle.com:3306)/mysql
	connURL := "root:en98as@tcp(db.mabetle.com:3306)/mysql"
	PingDB(connURL)
}
