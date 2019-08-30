package http

import (
	"io/ioutil"
	_http "net/http"
	"strings"
)

//IP get ip from http request
func IP(r *_http.Request) string {
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

//URL2Bytes get content bytes from a url
func URL2Bytes(url string) ([]byte, error) {
	resp, err := _http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//URL2String get content string from a url
func URL2String(url string) (string, error) {
	body, err := URL2Bytes(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
