package main

import (
	"flag"
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	. "net/http"
	"strings"
)

func main() {

	const (
		defaultPort               = ":8888"
		defaultPortUsage          = "default server port, ':8888', ':8080', ':80'..."
		defaultHost               = "127.0.0.1:3000"
		defaultHostUsage          = "default target host, '127.0.0.1:3000'"
		defaultExtendHeaders      = "GoProxy:go-proxy"
		defaultExtendHeadersUsage = "default extend header, 'GoProxy:go-proxy"
		defaultVerbose            = false
		defaultVerboseUsage       = "default verbose, 'false'"
	)

	// flags
	port := flag.String("port", defaultPort, defaultPortUsage)
	host := flag.String("host", defaultHost, defaultHostUsage)
	extendheaders := flag.String("extend-headers", defaultExtendHeaders, defaultExtendHeadersUsage)
	verbose := flag.Bool("verbos", defaultVerbose, defaultVerboseUsage)

	flag.Parse()

	fmt.Println("server will run on : %s", *port)
	fmt.Println("override proxy host on :%s", *host)
	fmt.Println("extend headers on :%s", *extendheaders)

	// proxy
	proxy := goproxy.NewProxyHttpServer()
	if *verbose == true {
		proxy.Verbose = true
	}

	proxy.OnRequest(goproxy.DstHostIs(*host)).DoFunc(
		func(r *Request, ctx *goproxy.ProxyCtx) (*Request, *Response) {
			slice := strings.Split(*extendheaders, ",")
			for _, extendheader := range slice {
				fmt.Println("extend header : %s", extendheader)
				header := strings.Split(extendheader, ":")
				r.Header.Set(header[0], header[1])
			}
			return r, nil
		})
	log.Fatal(ListenAndServe(*port, proxy))
}
