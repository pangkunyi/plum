package files

import (
	"testing"
)

func TestScanFile(t *testing.T) {
	ScanFile("/tmp/a.log", func(line string) error {
		t.Log(line)
		return nil
	})
}
