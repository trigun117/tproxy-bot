package code

import (
	"testing"
)

func TestGetProxies(t *testing.T) {
	if err := GetProxies(); err != nil {
		t.Fail()
	}
}
