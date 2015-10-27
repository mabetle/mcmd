package main

import (
	"flag"
	"fmt"
	"os"
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

var(
	username string
	passwd string
	port int
	host string
	dbname string
)


func doFlag(){
	flag.StringVar(&username,"u","demo","db username")
	flag.StringVar(&passwd,"p","demo","db password")
	flag.IntVar(&port,"port",3306,"db port")
	flag.StringVar(&host,"host","127.0.0.1", "db host")

	flag.Usage=func(){
		fmt.Fprintf(os.Stderr,"Usage: %s [-u -p -h -port] dbname", os.Args[0])
		flag.PrintDefaults()
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

	// connURL format: xxx:xxx@tcp(ip:3306)/xxx
	connURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",username,passwd,host,port,dbname)

	PingDB(connURL)
}
