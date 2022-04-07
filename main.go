package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/jthomperoo/simple-proxy/proxy"
)

var (
	Version = "development"
)

const (
	httpProtocol  = "http"
	httpsProtocol = "https"
)

func init() {
	flag.Set("logtostderr", "true")
}

func main() {
	var version bool
	flag.BoolVar(&version, "version", false, "prints current simple-proxy version")
	var protocol string
	flag.StringVar(&protocol, "protocol", httpProtocol, "proxy protocol (http or https)")
	var port string
	flag.StringVar(&port, "port", "8888", "proxy port to listen on")
	var certPath string
	flag.StringVar(&certPath, "cert", "", "path to cert file")
	var keyPath string
	flag.StringVar(&keyPath, "key", "", "path to key file")
	var timeoutSecs int
	flag.IntVar(&timeoutSecs, "timeout", 10, "timeout in seconds")
	flag.Parse()

	if version {
		fmt.Println(Version)
		os.Exit(0)
	}

	if protocol != httpProtocol && protocol != httpsProtocol {
		glog.Fatalln("Protocol must be either http or https")
	}

	if protocol == httpsProtocol && (certPath == "" || keyPath == "") {
		glog.Fatalf("If using HTTPS protocol --cert and --key are required")
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: proxy.NewProxyHandler(timeoutSecs),
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	if protocol == httpProtocol {
		glog.V(0).Infoln("Starting HTTP proxy...")
		log.Fatal(server.ListenAndServe())
	} else {
		glog.V(0).Infoln("Starting HTTPS proxy...")
		log.Fatal(server.ListenAndServeTLS(certPath, keyPath))
	}
}
