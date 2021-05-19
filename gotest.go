package gotest

import (
	"strconv"

	"golang.org/x/text/date"
)

func SomeTest() string {
	return strconv.Itoa(date.EtcUTC)
}
