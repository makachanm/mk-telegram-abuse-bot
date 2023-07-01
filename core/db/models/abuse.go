package dbmodels

import (
	"abusebot/core/db"
	"fmt"
	"strings"
)

type AbuseModel struct {
	DB *db.DatabaseCore
}

func NewAbuseModel(dbc *db.DatabaseCore) AbuseModel {
	return AbuseModel{DB: dbc}
}

func (am *AbuseModel) InsertAbuse(AbuseIDs []string) {
	var abids []string = make([]string, 0)

	for _, val := range AbuseIDs {
		abids = append(abids, (`('` + val + `')`))
	}

	var abjoin string

	if len(abids) >= 1 {
		abjoin = strings.Join(abids, ", ")
	} else {
		abjoin = abids[0]
	}

	query := fmt.Sprint("INSERT INTO abuseIDs (ab_id) VALUES ", abjoin)

	qo, err := am.DB.ExcuteQuery(query)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DB INSERTED ACTION: ", qo)
}

func (am *AbuseModel) DeleteAbuse(AbuseIDs []string) {
	var abids []string = make([]string, 0)

	for _, val := range AbuseIDs {
		abids = append(abids, (`'` + val + `'`))
	}

	var abjoin string

	if len(abids) >= 1 {
		abjoin = strings.Join(abids, ", ")
	} else {
		abjoin = abids[0]
	}

	query := fmt.Sprint("DELETE FROM abuseIDs WHERE ab_id IN (", abjoin, ")")

	qo, err := am.DB.ExcuteQuery(query)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DB DELETED ACTION: ", qo)
}

func (am *AbuseModel) GetAbuse() []string {
	var result []string = make([]string, 0)

	query := "SELECT ab_id FROM abuseIDs"

	rows, err := am.DB.GetDataFromQuery(query)
	if err != nil {
		fmt.Println(err)
		rows.Close()
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
