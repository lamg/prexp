package main

import (
	"context"
	"flag"
	"fmt"
	gp "github.com/lamg/goproxy"
	"log"
	"net"
	"net/http"
)

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8080", "proxy listen address")
	flag.Parse()
	px := gp.NewProxyHttpServer()
	px.Verbose = *verbose
	px.Tr.DialContext = dialHTTP
	log.Fatal(http.ListenAndServe(*addr, px))
}

type proxy struct {
	px *gp.ProxyHttpServer
}

// RemoteAddr is the type to be used as key
// of RemoteAddr value in context
type RemoteAddr string

func (p *proxy) serveHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.WithContext(context.WithValue(context.Background(),
		RemoteAddr("RemoteAddress"), r.RemoteAddr))
	p.serveHTTP(w, q)
}

func dialHTTP(c context.Context, nt, ad string) (n net.Conn, e error) {
	v := c.Value(RemoteAddr("RemoteAddress"))
	fmt.Printf("%v\n", v)
	n, e = net.Dial(nt, ad)
	return
}
