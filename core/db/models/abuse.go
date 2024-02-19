package dbmodels

import (
	"abusebot/core/db"
	"fmt"
)

type AbuseModel struct {
	DB *db.DatabaseCore
}

type AbuseIDCtx struct {
	AbuseID string `db:"ab_id"`
}

func NewAbuseModel(dbc *db.DatabaseCore) AbuseModel {
	return AbuseModel{DB: dbc}
}

func (am *AbuseModel) InsertAbuse(AbuseIDs []string) {
	var abids []AbuseIDCtx = make([]AbuseIDCtx, 0)

	for _, val := range AbuseIDs {
		abids = append(abids, AbuseIDCtx{AbuseID: val})
	}

	qo, err := am.DB.DBDriver.NamedExec("INSERT INTO abuseIDs (ab_id) VALUES (:ab_id)", abids)
	if err != nil {
		fmt.Println(err)
	}

	affectrow, _ := qo.RowsAffected()

	fmt.Println("DB INSERTED ACTION: ", affectrow, "AFFECTED")
}

func (am *AbuseModel) DeleteAbuse(AbuseIDs []string) {
	var abids []AbuseIDCtx = make([]AbuseIDCtx, 0)

	for _, val := range AbuseIDs {
		abids = append(abids, AbuseIDCtx{AbuseID: val})
	}

	//query := fmt.Sprint("DELETE FROM abuseIDs WHERE ab_id IN (", abjoin, ")")

	qo, err := am.DB.DBDriver.NamedExec("DELETE FROM abuseIDs WHERE ab_id IN (:ab_id)", abids)
	if err != nil {
		fmt.Println(err)
	}

	affectrow, _ := qo.RowsAffected()

	fmt.Println("DB DELETED ACTION: ", affectrow, " AFFECTED")
}

func (am *AbuseModel) GetAbuse() []string {
	var result []string = make([]string, 0)

	qo, err := am.DB.DBDriver.Prepare("SELECT ab_id FROM abuseIDs")
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	rows, err := qo.Query()
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	if err != nil {
		fmt.Println(err)
		rows.Scan()
	}

	var abuse_d string

	for rows.Next() {
		rowerr := rows.Scan(&abuse_d)
		if rowerr != nil {
			fmt.Println("READING ROWS IN ERROR: ", rowerr)
		} else {
			result = append(result, abuse_d)
		}
	}
	defer rows.Close()

	fmt.Println("READING: ", result)

	return result
}
