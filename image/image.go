package image

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
)

//ToDataURIWithResize encode an url image and resizes it, then encode to base64 and make the data URI
func ToDataURIWithResize(url string, width, height int) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	contentType := strings.ToLower(resp.Header.Get("Content-Type"))
	img, err := imaging.Decode(resp.Body)
	if err != nil {
		return "", err
	}
	img = imaging.Resize(img, width, height, imaging.Lanczos)
	format := imaging.JPEG
	if contentType == "image/jpeg" {
		format = imaging.JPEG
	} else if contentType == "image/png" {
		format = imaging.PNG
	} else if contentType == "image/tiff" {
		format = imaging.TIFF
	} else if contentType == "image/bmp" {
		format = imaging.BMP
	} else if contentType == "image/gif" {
		format = imaging.GIF
	} else {
		return "", fmt.Errorf("unsupported image type:%s", contentType)
	}
	var buf bytes.Buffer
	err = imaging.Encode(&buf, img, format)
	if err != nil {
		return "", err
	}
	return "data:" + contentType + ";base64," + base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

//ToDataURI encode an url content to base64 and make the data URI
func ToDataURI(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	contentType := resp.Header.Get("Content-Type")
	return "data:" + contentType + ";base64," + base64.StdEncoding.EncodeToString(content), nil
}
