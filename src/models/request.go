package models

import (
	"time"
)

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
	CreatedAt     time.Time // `gorm:"->"`
}
