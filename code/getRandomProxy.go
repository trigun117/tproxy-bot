package code

import (
	"strings"
)

// GetRandomProxy return link with random proxy
func GetRandomProxy() (link string) {

	getProxies()

	// Get random proxy and put it to link
	randomNumber := random(0, len(p.Proxies))
	randomProxy := strings.Split(p.Proxies[randomNumber], ":")
	link = `http://t.me/socks?server=` + randomProxy[0] + `&port=` + randomProxy[1]

	return
}
