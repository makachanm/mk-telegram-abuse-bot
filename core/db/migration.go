package db

import (
	"github.com/jmoiron/sqlx"
)

type DBMigration struct {
	DB *sqlx.DB
}

func NewDBMigration(dp *sqlx.DB) DBMigration {
	return DBMigration{DB: dp}
}

const Scheme = `CREATE TABLE IF NOT EXISTS abuseIDs ( 
	ab_id varchar(255)
)`

func (dm *DBMigration) InitalizeTable() {
	_, err := dm.DB.Exec(Scheme)

	if err != nil {
		dm.DB.Close()
		panic(err)
	}
}
