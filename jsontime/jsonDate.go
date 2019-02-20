package jsontime

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	//"github.com/astaxie/beego/utils"
)

type JsonDate time.Time

func (p *JsonDate) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(formatDate, string(data), time.Local)
	*p = JsonDate(local)
	return err
}

func (c JsonDate) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0)
	data = append(data, '"')
	data = time.Time(c).AppendFormat(data, formatDate)
	data = append(data, '"')
	return data, nil
}

func (c JsonDate) String() string {
	return time.Time(c).Format(formatDate)
}

func (c JsonDate) Value() time.Time {
	return time.Time(c)
}
func (e *JsonDate) Set(d time.Time) {
	*e = JsonDate(d)
}

/*
// String convert time to string
func (e *JsonTime) String() string {
	return e.Value().String()
}
*/

// FieldType return enum type Date
func (e *JsonDate) FieldType() int {
	return orm.TypeDateField
}

// SetRaw convert the interface to time.Time. Allow string and time.Time
func (e *JsonDate) SetRaw(value interface{}) error {
	switch d := value.(type) {
	case time.Time:
		e.Set(d)
	case string:
		v, err := timeParse(d, formatDateTime)
		if err == nil {
			e.Set(v)
		}
		return err
	default:
		return fmt.Errorf("<JsonTime.SetRaw> unknown value `%s`", value)
	}
	return nil
}

// RawValue return time value
func (e *JsonDate) RawValue() interface{} {
	return e.Value()
}

var _ orm.Fielder = new(JsonDate)
