package comm

import (
	"strings"
	"time"
)

/*
필요한 데이터 정리

시간 계산에 필요한 것 정리하기


new Date => string
: yyyy-mm-dd
: yyyy-mm-dd HH:ii:ss

string => new Date

string => unixtime

시간 비교하기

string,string => int  // unixtime 으로 바꿔서 second로 계산하기

string , int => string  // 몇일 차이 를 계산하기

string , int => date //

*/

// Date(thisDate string)
func Time(c time.Time) string {
	timeT, _ := time.Parse("2006-01-02", Format("yyyy-mm-dd hh:ii:ss", c))
	return Format("hh:ii:ss", timeT)
}

// string => time.Time
func DateTime(thisDate string) time.Time {
	var sendTime time.Time
	if len(thisDate) <= 10 {
		sendTime, _ = time.Parse("2006-01-02", thisDate)
	} else {
		sendTime, _ = time.Parse("2006-01-02 15:04:05", thisDate)
	}
	return sendTime
}

// time.Time => string format
func Format(types string, thisDate time.Time) string {
	formatTypes := makeFormat(types)
	return thisDate.Format(formatTypes)
}

func makeFormat(format string) string {
	returnstr := format
	returnstr = changeString(returnstr, "yyyy", "2006")
	returnstr = changeString(returnstr, "mm", "01")
	returnstr = changeString(returnstr, "dd", "02")
	returnstr = changeString(returnstr, "hh", "15")
	returnstr = changeString(returnstr, "ii", "04")
	returnstr = changeString(returnstr, "ss", "05")
	return returnstr
}

// origin find, change
func changeString(origin, find, change string) string {
	if strings.Contains(origin, find) {
		origin = strings.Replace(origin, find, change, 1)
	}
	return origin
}
