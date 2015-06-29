package http

import (
	_http "net/http"
	"strings"
)

func Ip(r _http.Request) string {
	address := r.Header.Get("X-Real-IP")
	if address != "" && address != "unknown" {
		return address
	}
	address = r.Header.Get("X-Forwarded-For")
	if address != "" && address != "unknown" {
		return address
	}
	address = r.Header.Get("Proxy-Client-IP")
	if address != "" && address != "unknown" {
		return address
	}
	address = r.Header.Get("WL-Proxy-Client-IP")
	if address != "" && address != "unknown" {
		return address
	}
	address = r.RemoteAddr
	return address[:strings.Index(address, ":")]
}
