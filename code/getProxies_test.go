package code

import (
	"os"
	"testing"
)

func TestGetProxies(t *testing.T) {
	os.Setenv("URL", "http://tproxyt.tk/json")
	os.Setenv("FI", "password")
	if err := getProxies(); err != nil {
		t.Fail()
	}
}
