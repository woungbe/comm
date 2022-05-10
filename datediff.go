package comm

import (
	"reflect"
	"time"
)

// 날짜 비교 계산 관련 데이터 처리
// start, end  =>  return end - start
func DiffDate(aa interface{}, bb interface{}) interface{} {
	return separate(aa, bb)
}

func AddDate(aa interface{}, bb interface{}) interface{} {
	return addDates(aa, bb.(int))
}

func AddSecond(aa interface{}, bb interface{}) interface{} {
	return addTime(aa, bb.(int))
}

func separate(aa, bb interface{}) interface{} {
	if CheckTypes(aa) == "string" && CheckTypes(bb) == "string" {
		// 날짜끼리 비교해서 빼기
		return stringNstring(aa.(string), bb.(string))
	}

	return nil
}

// string,string 으로 해서 무조건 second로 나타내줌
func stringNstring(aa, bb string) int {
	t1 := DateTime(aa)
	t2 := DateTime(bb)
	second := t2.Sub(t1).Seconds()
	// days
	return int(second)
}

// date, second => dateTime
func addDates(aa interface{}, bb int) time.Time {
	tmp := CheckDiv(aa, bb)
	t1 := DateTime(tmp)
	return t1.AddDate(0, 0, bb)
}

// string, second
func addTime(aa interface{}, bb int) time.Time {
	tmp := CheckDiv(aa, bb)

	hour := bb / 60 / 60
	minute := bb / 60 % 60
	second := bb % 60 % 60
	t1 := DateTime(tmp).Local().Add(
		time.Hour*time.Duration(hour) +
			time.Minute*time.Duration(minute) +
			time.Second*time.Duration(second))
	return t1
}

func CheckDiv(aa interface{}, bb int) string {
	var tmp string
	if CheckTypes(aa) == "string" {
		tmp = aa.(string)
	} else if CheckTypes(aa) == "time.Time" {
		tmp = Format("yyyy-mm-dd hh:ii:ss", aa.(time.Time))
	}

	return tmp
}

func CheckTypes(aa interface{}) string {
	return reflect.TypeOf(aa).String()
}
