package proxy

import (
	"io"
	"net"
	"net/http"
	"time"

	"github.com/golang/glog"
)

func NewProxyHandler(timeoutSeconds int) *ProxyHandler {
	return &ProxyHandler{
		Timeout: time.Duration(timeoutSeconds) * time.Second,
	}
}

type ProxyHandler struct {
	Timeout time.Duration
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	glog.V(1).Infof("Serving '%s' request from '%s' to '%s'\n", r.Method, r.RemoteAddr, r.Host)
	if r.Method == http.MethodConnect {
		handleTunneling(w, r, p.Timeout)
	} else {
		handleHTTP(w, r)
	}
}

func handleTunneling(w http.ResponseWriter, r *http.Request, timeout time.Duration) {
	dest_conn, err := net.DialTimeout("tcp", r.Host, timeout)
	if err != nil {
		glog.Errorf("Failed to dial host, %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		glog.Errorln("Attempted to hijack connection that does not support it")
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	client_conn, _, err := hijacker.Hijack()
	if err != nil {
		glog.Errorf("Failed to hijack connection, %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	go transfer(dest_conn, client_conn)
	go transfer(client_conn, dest_conn)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

func handleHTTP(w http.ResponseWriter, req *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		glog.Errorf("Failed to proxy request, %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
