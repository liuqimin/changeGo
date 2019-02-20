package jsontime

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
	//"github.com/astaxie/beego/utils"
)

type JsonTime time.Time

func (p *JsonTime) UnmarshalJSON(data []byte) error {
	fmt.Printf("%q 2222", string(data))
	fort := strings.Trim(string(data), "\"")
	//*(*string)(unsafe.Pointer(&data))
	//fort := strings.Replace(string(data),`\"`,"",-1)
	//fort := fmt.Sprintf("\"%s\"", string(data))
	fmt.Printf("%q 2222", string(fort))
	local, err := time.ParseInLocation(formatDateTime, fort, time.Local)
	*p = JsonTime(local)
	if err != nil {
		fmt.Println("??err")
		fmt.Println("%v aaa", err)
	}
	return err
}

func (c JsonTime) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0)
	data = append(data, '"')
	data = time.Time(c).AppendFormat(data, formatDateTime)
	data = append(data, '"')
	return data, nil
}

func (c JsonTime) String() string {
	return time.Time(c).Format(formatDateTime)
}

func Todate(in string) (out time.Time, err error) {
	out, err = time.Parse(formatDate, in)
	return out, err
}

func Todatetime(in string) (out time.Time, err error) {
	out, err = time.Parse(formatDateTime, in)
	return out, err
}

func (c JsonTime) Value() time.Time {
	return time.Time(c)
}
func (e *JsonTime) Set(d time.Time) {
	*e = JsonTime(d)
}

/*
// String convert time to string
func (e *JsonTime) String() string {
	return e.Value().String()
}
*/

// FieldType return enum type Date
func (e *JsonTime) FieldType() int {
	return orm.TypeDateField
}

// SetRaw convert the interface to time.Time. Allow string and time.Time
func (e *JsonTime) SetRaw(value interface{}) error {
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
func (e *JsonTime) RawValue() interface{} {
	return e.Value()
}

var _ orm.Fielder = new(JsonTime)
