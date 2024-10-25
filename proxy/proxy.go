package proxy

import (
	"encoding/base64"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/golang/glog"
	net_proxy "golang.org/x/net/proxy"
)

// Using a socks5 proxy to tunnel the HTTP requests, e.g., 127.0.0.1:7890
var Socks5 = ""

func NewProxyHandler(timeoutSeconds int) *ProxyHandler {
	return &ProxyHandler{
		Timeout: time.Duration(timeoutSeconds) * time.Second,
	}
}

type ProxyHandler struct {
	Timeout    time.Duration
	Username   *string
	Password   *string
	LogAuth    bool
	LogHeaders bool
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	glog.V(1).Infof("Serving '%s' request from '%s' to '%s'\n", r.Method, r.RemoteAddr, r.Host)
	if p.LogHeaders {
		for name, values := range r.Header {
			for i, value := range values {
				glog.V(1).Infof("'%s': [%d] %s", name, i, value)
			}
		}
	}
	if p.Username != nil && p.Password != nil {
		username, password, ok := proxyBasicAuth(r)
		if !ok || username != *p.Username || password != *p.Password {
			if p.LogAuth {
				glog.Errorf("Unauthorized, username: %s, password: %s\n", username, password)
			} else {
				glog.Errorln("Unauthorized")
			}
			w.Header().Set("Proxy-Authenticate", "Basic")
			http.Error(w, "Unauthorized", http.StatusProxyAuthRequired)
			return
		}
	}
	if r.Method == http.MethodConnect {
		handleTunneling(w, r, p.Timeout)
	} else {
		handleHTTP(w, r)
	}
}

func handleTunneling(w http.ResponseWriter, r *http.Request, timeout time.Duration) {
	var dest_conn net.Conn
	var err error

	// Check if the socks5 proxy is set
	// Then tunnel to socks5 proxy
	if Socks5 == "" {
		dest_conn, err = net.DialTimeout("tcp", r.Host, timeout)
	} else {
		var socks5_dailer net_proxy.Dialer
		socks5_dailer, err = net_proxy.SOCKS5("tcp", Socks5, nil, &net.Dialer{
			Timeout:   timeout,
			KeepAlive: 30 * time.Second,
		})
		if err != nil {
			glog.Errorf("Failed to dail socks5 proxy %s, %s\n", Socks5, err.Error())
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		dest_conn, err = socks5_dailer.Dial("tcp", r.Host)
	}
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

func proxyBasicAuth(r *http.Request) (username, password string, ok bool) {
	auth := r.Header.Get("Proxy-Authorization")
	if auth == "" {
		return
	}
	return parseBasicAuth(auth)
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	// Case insensitive prefix match. See Issue 22736.
	if len(auth) < len(prefix) || !equalFold(auth[:len(prefix)], prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

// EqualFold is strings.EqualFold, ASCII only. It reports whether s and t
// are equal, ASCII-case-insensitively.
func equalFold(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if lower(s[i]) != lower(t[i]) {
			return false
		}
	}
	return true
}

// lower returns the ASCII lowercase version of b.
func lower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
