package http

import (
	"io/ioutil"
	_http "net/http"
	"strings"
)

func Ip(r *_http.Request) string {
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

func Url2Bytes(url string) ([]byte, error) {
	resp, err := _http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func Url2String(url string) (string, error) {
	body, err := Url2Bytes(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
