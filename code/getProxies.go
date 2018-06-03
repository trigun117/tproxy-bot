package code

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// P store unmarshaled proxies
var p Proxy

// Proxy is a struct which contain proxies
type Proxy struct {
	Proxies []string
}

var (
	link     = os.Getenv("URL")
	password = os.Getenv("PASS")
	field    = os.Getenv("FI")
)

// GetProxies fetch json with proxies
func getProxies() (err error) {

	re, _ := http.PostForm(link, url.Values{field: {password}})

	//read bytes
	b, _ := ioutil.ReadAll(re.Body)

	//unmarshal json
	json.Unmarshal(b, &p)

	re.Body.Close()
	return
}
