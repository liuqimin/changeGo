package jsontime

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	formatTime     = "15:04:05"
	formatDate     = "2006-01-02"
	formatDateTime = "2006-01-02 15:04:05"
)

func timeParse(dateString, format string) (time.Time, error) {
	tp, err := time.ParseInLocation(format, dateString, orm.DefaultTimeLoc)
	return tp, err
}
