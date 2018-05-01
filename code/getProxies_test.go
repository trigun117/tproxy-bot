package code

import (
	"testing"
)

func TestGetProxies(t *testing.T) {
	if err := getProxies(); err != nil {
		t.Fail()
	}
}
