package code

import (
	"testing"
)

func TestGetRandomProxy(t *testing.T) {
	if p := GetRandomProxy(); p == "" {
		t.Fail()
	}
}
