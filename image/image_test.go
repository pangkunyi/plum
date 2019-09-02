package image

import (
	"strings"
	"testing"
)

func TestToDataURIWithResize(t *testing.T) {
	content, err := ToDataURIWithResize("https://www.baidu.com/img/superlogo_c4d7df0a003d3db9b65e9ef0fe6da1ec.png?where=super", 100, 100)
	if err != nil {
		t.Errorf("failed to encode image to base64, cause by:%v", err)
		return
	}
	if !strings.HasPrefix(content, "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGQAAABkCAYAAABw4pVUAAAZ7ElEQVR4nOx8C5QUxbl/VXX37MzsY/YBuAqLCITHXxd0lRhAzCHR+AgGUP4aERUSH1H") {
		t.Errorf("failed to encode, content:%s", content)
	}
}

func TestToDataURI(t *testing.T) {
	content, err := ToDataURI("https://www.baidu.com/img/superlogo_c4d7df0a003d3db9b65e9ef0fe6da1ec.png?where=super")
	if err != nil {
		t.Errorf("failed to encode image to base64, cause by:%v", err)
		return
	}
	if !strings.HasPrefix(content, "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAhwAAAECCAMAAACCFP44AAACPVBMVEUAAAAGAgIFAgIEAAAEAAAEAAAEAAAEAQEGAgIEAAAEAAAEAQEEAQH") {
		t.Errorf("failed to encode, content:%s", content)
	}
}
