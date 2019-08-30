package image

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

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
