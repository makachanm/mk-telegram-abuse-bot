package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseCore struct {
	DBDriver *sqlx.DB
}

func NewDatabaseCore(db *sqlx.DB) DatabaseCore {
	return DatabaseCore{DBDriver: db}
}

func StartDB(dbfile string) *sqlx.DB {
	var serr error
	var sqldb *sqlx.DB

	//sqldb, serr = sql.Open("sqlite3", dbfile)
	sqldb, serr = sqlx.Open("sqlite3", dbfile)
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
