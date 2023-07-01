package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseCore struct {
	DBDriver *sql.DB
}

func NewDatabaseCore(db *sql.DB) DatabaseCore {
	return DatabaseCore{DBDriver: db}
}

func StartDB(dbfile string) *sql.DB {
	var serr error
	var sqldb *sql.DB

	sqldb, serr = sql.Open("sqlite3", dbfile)
	if serr != nil {
		log.Fatalf(serr.Error())
		sqldb.Close()
	}
	fmt.Println("SQLite3 Document Loaded.")

	fmt.Print("Checking SQLite3...")
	perr := sqldb.Ping()
	if perr != nil {
		log.Fatalf(perr.Error())
		sqldb.Close()
	}

	fmt.Println("OK.")

	fmt.Print("Excuting Migration...")
	migDB := NewDBMigration(sqldb)
	migDB.InitalizeTable()

	fmt.Println("OK.")

	return sqldb
}

func (dc *DatabaseCore) ExcuteQuery(query string) (sql.Result, error) {
	pr, perr := dc.DBDriver.Prepare(query)
	if perr != nil {
		log.Fatalf(perr.Error())
		pr.Close()
	}

	return pr.Exec()
}

func (dc *DatabaseCore) GetDataFromQuery(query string) (*sql.Rows, error) {
	qrw, qerr := dc.DBDriver.Query(query)
	fmt.Println(qerr)
	return qrw, qerr
}
