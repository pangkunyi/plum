package files

import (
	"testing"
)

func TestScanFile(t *testing.T) {
	ScanFile("/tmp/a.log", func(line string) error {
		//t.Log(line)
		t.Logf("line size:%v\n", len(line))
		return nil
	})
	ScanFileFull("/tmp/a.log", func(line string) error {
		//t.Log(line)
		t.Logf("line size:%v\n", len(line))
		return nil
	})
}
