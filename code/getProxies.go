package code

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//P store unmarshaled proxies
var p Proxy

//Proxy is a struct which contain proxies
type Proxy struct {
	Proxies []string
}

//GetProxies fetch json with proxies
func getProxies() (err error) {

	re, _ := http.Get(`http://telegram-socks.tk/json`)

	//read bytes
	b, _ := ioutil.ReadAll(re.Body)

	//unmarshal json
	json.Unmarshal(b, &p)

	re.Body.Close()
	return
}
