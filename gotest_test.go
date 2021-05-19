package gotest

import (
	"fmt"
	"testing"

	"golang.org/x/text/date"
)

func TestSomeTest(t *testing.T) {
	want := fmt.Sprintf("%d", date.EtcUTC)
	if got := SomeTest(); got != want {
		t.Error(got, "not equal", want)
	}
}
