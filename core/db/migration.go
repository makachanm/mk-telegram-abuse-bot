package db

import "database/sql"

type DBMigration struct {
	DB *sql.DB
}

func NewDBMigration(dp *sql.DB) DBMigration {
	return DBMigration{DB: dp}
}

func (dm *DBMigration) InitalizeTable() {
	initQuery := `
	CREATE TABLE IF NOT EXISTS abuseIDs ( 
		ab_id varchar(255)
	)`
	_, err := dm.DB.Exec(initQuery)

	if err != nil {
		dm.DB.Close()
		panic(err)
	}
	//defer dm.DB.Close()
}
