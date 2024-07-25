package main

import (
	"fmt"
	"net"
	neturl "net/url"
)

func main() {
	pln := fmt.Println
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	url, err := neturl.Parse(s)
	if err != nil {
		panic(err)
	}

	pln(url.Scheme)
	pln(url.User)
	pln(url.User.Username())
	p, _ := url.User.Password()
	pln(p)

	pln(url.Host)
	host, port, _ := net.SplitHostPort(url.Host)
	pln(host, ":", port)

	pln(url.Path)
	pln(url.Fragment)

	pln(url.RawQuery)
	m, _ := neturl.ParseQuery(url.RawQuery)
	pln(m)
	pln(m["k"][0])
}
