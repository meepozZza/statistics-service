package models

import "time"

type Request struct {
	Model
	UserId        uint64
	RequestFrom   string
	RequestTo     string
	RequestMethod string
	RequestBody   string
	ResponseBody  string
	ResponseCode  uint32
	ResponseTime  int64
	CreatedAt     time.Time `gorm:"->"`
}

type Report struct {
	DAU    uint32 `db:"DAU"`
	MAU    uint32 `db:"MAU"`
	VD     uint32 `db:"VD"`
	Views  uint32 `db:"Views"`
	Visits uint32 `db:"Visits"`
	CRtM   uint32 `db:"CRtM"`
	UC     uint32 `db:"UC"`
	AC     uint32 `db:"AC"`
	RT     uint32 `db:"RT"`
}
