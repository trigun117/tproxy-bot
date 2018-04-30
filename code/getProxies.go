package code

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//P store unmarshaled proxies
var P Proxy

//Proxy is a struct which contain proxies
type Proxy struct {
	Proxies []string
}

//GetProxies fetch json with proxies
func GetProxies() (err error) {

	re, err := http.Get(`http://telegram-socks.tk/json`)
	if err != nil {
		return
	}

	defer re.Body.Close()

	//read bytes
	b, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return
	}

	//unmarshal json
	json.Unmarshal(b, &P)
	return
}
