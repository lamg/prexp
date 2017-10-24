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

func dialHTTP(c context.Context, nt, ad string) (n net.Conn, e error) {
	fmt.Printf("%v\n", c)
	n, e = net.Dial(nt, ad)
	return
}
