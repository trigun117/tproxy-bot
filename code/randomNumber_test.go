package code

import (
	"reflect"
	"testing"
)

func TestRandomNumber(t *testing.T) {
	if r := Random(3, 6); reflect.TypeOf(r).Kind() != reflect.Int {
		t.Fail()
	}
}
