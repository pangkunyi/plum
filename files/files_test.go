package files

import (
	"testing"
	"time"
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
	fl, err := NewFileLoader("/tmp/a.log", func(lines []string) (interface{}, error) {
		ls := make([]string, 0)
		for _, line := range lines {
			ls = append(ls, line)
		}
		return ls, nil
	})
	if err != nil {
		t.Fatalf("failed load file loader, cause by %s", err)
	}

	for i := 0; i < 5; i++ {
		for _, line := range fl.Value().([]string) {
			t.Logf("line:%v\n", line)
		}
		time.Sleep(5 * time.Second)
	}
}
