package services

import (
	"github.com/meepozZza/statistics-service/src/database"

	"github.com/blockloop/scan"
)

type Report struct {
	DAU    []DateValue
	MAU    []DateValue
	VD     []VD
	Views  []UserValue
	Visits []UserValue
	CRtM   uint32   `db:"CRtM"`
	UC     uint32   `db:"UC"`
	AC     *uint64  `db:"value"`
	RT     *float64 `db:"value"`
}

type DateValue struct {
	Value *uint64 `db:"value"`
	Date  *string `db:"date"`
}

type VD struct {
	Url    *string `db:"url"`
	UserId *uint64 `db:"user_id"`
	Value  *uint32 `db:"value"`
	Date   *string `db:"date"`
}

type UserValue struct {
	UserId *uint64 `db:"user_id"`
	Value  *uint64 `db:"value"`
}

func CalculateDAU() []DateValue {
	var data []DateValue = []DateValue{}

	rows, err := database.SqlDB.Query(`
		select count(distinct user_id) as value, toString(toDate(created_at)) as date
		from requests
		group by date
	`)

	if err != nil {
		return data
	}

	if err = scan.Rows(&data, rows); err != nil {
		return data
	}

	return data
}

func CalculateMAU() []DateValue {
	var data []DateValue = []DateValue{}

	rows, err := database.SqlDB.Query(`
		select count(distinct user_id) as value, toString(toStartOfMonth(created_at)) as date
		from requests
		group by date
	`)

	if err != nil {
		return data
	}

	if err = scan.Rows(&data, rows); err != nil {
		return data
	}

	return data
}

func CalculateVD() []VD {
	var data []VD = []VD{}

	rows, err := database.SqlDB.Query(`
		select request_to as url, user_id, (if(neighbor(user_id, 1) = user_id AND neighbor(request_to, 1) = request_to, neighbor(created_at, 1) - created_at, 0)) as value, toString(created_at) as date
		from requests 
		where isNotNull(url) and isNotNull(user_id) 
		order by user_id, created_at
	`)

	if err != nil {
		return data
	}

	if err = scan.Rows(&data, rows); err != nil {
		return data
	}

	return data
}

func CalculateViews() []UserValue {
	var data []UserValue = []UserValue{}

	rows, err := database.SqlDB.Query(`
		select user_id, count() as value
		from requests
		group by user_id
	`)

	if err != nil {
		return data
	}

	if err = scan.Rows(&data, rows); err != nil {
		return data
	}

	return data
}

func CalculateVisits() []UserValue {
	var data []UserValue = []UserValue{}

	rows, err := database.SqlDB.Query(`
		select
			user_id,
			sum(
				if (
					neighbor(r.user_id, -1, 0) != r.user_id,
					1,
					dateDiff('second', neighbor(r.created_at, -1), r.created_at) > 60 * 30
				)
			) as value
		from (
			select *
			from requests
			order by user_id, created_at
		) as r
		group by user_id
	`)

	if err != nil {
		return data
	}

	if err = scan.Rows(&data, rows); err != nil {
		return data
	}

	return data
}

func CalculateAC() *uint64 {
	var data *uint64

	rows, err := database.SqlDB.Query(`
		select count()
		from requests
	`)

	if err != nil {
		return data
	}

	if err = scan.Row(&data, rows); err != nil {
		return data
	}

	return data
}

func CalculateRT() *float64 {
	var data *float64

	rows, err := database.SqlDB.Query(`
		select sum(response_time) / count() as value
		from requests
	`)

	if err != nil {
		return data
	}

	if err = scan.Row(&data, rows); err != nil {
		return data
	}

	return data
}
