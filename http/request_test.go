package http

import (
	_http "net/http"
	"testing"
)

func TestIp(t *testing.T) {
	r := _http.Request{}
	r.RemoteAddr = "192.16.8.8:56781"
	if Ip(r) != "192.16.8.8" {
		t.Errorf("failed:%s", Ip(r))
	}

	r.Header = make(_http.Header)
	r.Header.Add("X-Real-IP", "10.2.2.2, 1.2.3.3")
	if Ip(r) != "10.2.2.2, 1.2.3.3" {
		t.Errorf("failed:%s", Ip(r))
	}
	r.Header.Del("X-Real-IP")
	r.Header.Add("X-Forwarded-For", "10.2.2.2, 1.2.3.4")
	if Ip(r) != "10.2.2.2, 1.2.3.4" {
		t.Errorf("failed:%s", Ip(r))
	}
	r.Header.Del("X-Forwarded-For")
	r.Header.Add("Proxy-Client-IP", "10.2.2.2, 1.2.3.5")
	if Ip(r) != "10.2.2.2, 1.2.3.5" {
		t.Errorf("failed:%s", Ip(r))
	}
	r.Header.Del("Proxy-Client-IP")
	r.Header.Add("WL-Proxy-Client-IP", "10.2.2.2, 1.2.3.6")
	if Ip(r) != "10.2.2.2, 1.2.3.6" {
		t.Errorf("failed:%s", Ip(r))
	}
}
