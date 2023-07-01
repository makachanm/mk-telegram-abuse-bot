package main

import (
	"abusebot/core"
	"abusebot/core/db"
	"abusebot/utils"
)

func main() {
	config, cerr := utils.LoadConfig()
	if cerr != nil {
		panic(cerr)
	}

	dCore := db.StartDB(config.DBName)
	database := db.NewDatabaseCore(dCore)

	service := core.NewService(&database, config)
	service.InitMainService()
}
