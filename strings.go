package comm

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func String(aa interface{}) (string, error) {
	attrTypes := reflect.TypeOf(aa).String()
	switch attrTypes {
	case "int":
		return strconv.Itoa(aa.(int)), nil
	case "int64":
		return strconv.FormatInt(aa.(int64), 10), nil
	case "float64":
		s := strconv.FormatFloat(aa.(float64), 'G', -1, 64)
		return s, nil
	}
	return "", nil
}

func Int(aa interface{}) (int, error) {
	attrTypes := reflect.TypeOf(aa).String()
	switch attrTypes {
	case "string":
		t1 := aa.(string)
		if IsComma(t1) {
			t1 = removeComma(t1)
		}

		if IsDot(t1) {
			t1 = removeDot(t1)
		}

		sv, _ := strconv.Atoi(t1)
		return sv, nil
	case "int64":
		return int(aa.(int64)), nil
	case "float64":
		fmt.Println("int to float64 ", int(aa.(float64)))
		tmps := int(aa.(float64))
		fmt.Println("tmps :", reflect.TypeOf(tmps))

		return int(aa.(float64)), nil
	}
	return 0, nil
}

func Int64(aa interface{}) (int64, error) {
	attrTypes := reflect.TypeOf(aa).String()
	switch attrTypes {
	case "string":
		t1 := CheckInt64(aa.(string))
		i, err := strconv.ParseInt(t1, 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	case "int":
		return int64(aa.(int)), nil
	case "float64":
		return int64(aa.(float64)), nil
	}
	return 0, nil
}

func Float64(aa interface{}) (float64, error) {
	attrTypes := reflect.TypeOf(aa).String()
	switch attrTypes {
	case "string":
		s := aa.(string)
		if IsComma(s) {
			s = removeComma(s)
		}

		i, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}

		return i, nil
	case "int":
		return float64(aa.(int)), nil
	case "int64":
		return float64(aa.(int64)), nil
	}
	return 0, nil
}

func CheckInt64(aa string) string {

	// 컴마가 있으면 일단 날려주구
	if IsComma(aa) {
		aa = removeComma(aa)
	}

	// 점이 있으면 없애주구
	if IsDot(aa) {
		aa = removeDot(aa)
	}

	return aa

}

// 컴파가 있나 없나 ?
func IsComma(aa string) bool {
	return strings.Contains(aa, ",")
}

// 점이 있나?
func IsDot(aa string) bool {
	return strings.Contains(aa, ".")
}

func removeComma(aa string) string {
	return strings.Replace(aa, ",", "", 10)
}

func removeDot(aa string) string {
	if IsDot(aa) {
		tmp1, _ := Float64(aa)
		tmp2 := math.Floor(tmp1)
		send, _ := String(tmp2)
		return send
	} else {
		return aa
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
