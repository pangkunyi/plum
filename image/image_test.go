package image

import (
	"strings"
	"testing"
)

func TestToDataURI(t *testing.T) {
	content, err := ToDataURI("https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo_top_86d58ae1.png")
	if err != nil {
		t.Errorf("failed to encode image to base64")
		return
	}
	if !strings.HasPrefix(content, "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAHU") {
		t.Errorf("failed to encode, content:%s", content)
	}
}
