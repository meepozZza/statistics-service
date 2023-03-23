package services

import (
	"statistics-service/src/database"

	"github.com/blockloop/scan"
)

func getDAU() uint32 {
	var DAU uint32 = 0

	rows, err := database.SqlDB.Query("SELECT t.DAU as DAU FROM (select count(distinct user_id) as DAU, toDate(created_at) as day from requests group by day ) as t LIMIT 1")

	if err != nil {
		return DAU
	}

	err = scan.Row(&DAU, rows)

	if err != nil {
		return DAU
	}

	return DAU
}
