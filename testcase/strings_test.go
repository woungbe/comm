package tests

import (
	"testing"

	"github.com/woungbe/comm"
)

// 주문값, 리턴값
func StringTestSample(t *testing.T, got interface{}, want string) {
	got, err := comm.String(got)
	if err != nil {
		t.Errorf("")
	}
	// want := "123123"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAdd(t *testing.T) {

	StringTestSample(t, 123123, "123123")
	StringTestSample(t, 123.123, "123.123")

	var aa int64
	aa = 123123
	StringTestSample(t, aa, "123123")

}

func StringTestInt(t *testing.T, got interface{}, want int) {
	got, err := comm.Int(got)
	if err != nil {
		t.Errorf(err.Error())
	}

	if got != want {
		t.Errorf("got %d, wanted %d", got.(int), want)
	}
}

func InterfaceToInt64Smaple(t *testing.T, got interface{}, want int64) {
	got1, err := comm.Int64(got)
	if err != nil {
		t.Errorf(err.Error())
	}

	if got1 != want {
		t.Errorf("got %d, wanted %d", got1, want)
	}
}

func InterfaceToFloat64Sample(t *testing.T, got interface{}, want float64) {
	got1, err := comm.Float64(got)
	if err != nil {
		t.Errorf(err.Error())
	}

	if got1 != want {
		t.Errorf("got %f, wanted %f", got1, want)
	}
}

func TestInt(t *testing.T) {

	StringTestInt(t, "123123", 123123)
	StringTestInt(t, "123.123", 123)
	StringTestInt(t, 123.123, 123)

	var aa int64
	aa = 123123
	StringTestInt(t, aa, 123123)

}

func TestInt64(t *testing.T) {
	var want1 int64 = 123123
	var want2 int64 = 123

	InterfaceToInt64Smaple(t, "123123", want1)
	InterfaceToInt64Smaple(t, "123,123", want1)
	InterfaceToInt64Smaple(t, "123.123", want2)
	InterfaceToInt64Smaple(t, 123.123, want2)
	InterfaceToInt64Smaple(t, 123, want2)
}

func TestFloat64(t *testing.T) {
	var want1 float64 = 1123.123
	var want2 float64 = 1123

	var got1 int = 1123
	var got2 int64 = 1123

	InterfaceToFloat64Sample(t, "1,123.123", want1)
	InterfaceToFloat64Sample(t, "1123.123", want1)
	InterfaceToFloat64Sample(t, got1, want2)
	InterfaceToFloat64Sample(t, got2, want2)
}
