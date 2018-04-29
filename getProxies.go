package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var p Proxy

//Proxy is a struct which contain proxies
type Proxy struct {
	Proxies []string
}

func getProxies() {
	re, _ := http.Get(`http://telegram-socks.tk/json`)
	b, _ := ioutil.ReadAll(re.Body)
	json.Unmarshal(b, &p)
	re.Body.Close()
}
