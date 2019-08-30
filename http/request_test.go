package http

import (
	_http "net/http"
	"testing"
)

func TestIP(t *testing.T) {
	r := &_http.Request{}
	r.RemoteAddr = "192.16.8.8:56781"
	if IP(r) != "192.16.8.8" {
		t.Errorf("failed:%s", IP(r))
	}

	r.Header = make(_http.Header)
	r.Header.Add("X-Real-IP", "10.2.2.2, 1.2.3.3")
	if IP(r) != "10.2.2.2, 1.2.3.3" {
		t.Errorf("failed:%s", IP(r))
	}
	r.Header.Del("X-Real-IP")
	r.Header.Add("X-Forwarded-For", "10.2.2.2, 1.2.3.4")
	if IP(r) != "10.2.2.2, 1.2.3.4" {
		t.Errorf("failed:%s", IP(r))
	}
	r.Header.Del("X-Forwarded-For")
	r.Header.Add("Proxy-Client-IP", "10.2.2.2, 1.2.3.5")
	if IP(r) != "10.2.2.2, 1.2.3.5" {
		t.Errorf("failed:%s", IP(r))
	}
	r.Header.Del("Proxy-Client-IP")
	r.Header.Add("WL-Proxy-Client-IP", "10.2.2.2, 1.2.3.6")
	if IP(r) != "10.2.2.2, 1.2.3.6" {
		t.Errorf("failed:%s", IP(r))
	}
}

func TestURL2String(t *testing.T) {
	url := "http://www.baidu.com"
	content, err := URL2String(url)
	if err != nil {
		t.Errorf("failed to get url:%s", url)
	}
	t.Log(content)
}
