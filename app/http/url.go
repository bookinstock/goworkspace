package http

import (
	"fmt"
	"net"
	"net/url"
)

func RunUrl() {

	p := fmt.Println

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, _ := url.Parse(s)

	p("scheme=", u.Scheme)
	p("user=", u.User)
	p("user.username=", u.User.Username())
	pass, _ := u.User.Password()
	p("user.password=", pass)
	p("host=", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	p("host=", host, "port", port)
	p("path=", u.Path)
	p("fragment=", u.Fragment)
	p("query=", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	p("parsed_query=", m)
	p("k=", m["k"][0])
}
