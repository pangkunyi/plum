package logs

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestLog(t *testing.T) {
	runtime.GOMAXPROCS(5)
	logger := NewLogger("/tmp/a.%s.log", true)
	var wg sync.WaitGroup
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				err := logger.Print(fmt.Sprintf("%d\x01sdkfjsdkfjs\x01ksdfjk\x01ksdfksdf\n", j*1000+i))
				if err != nil {
					t.Errorf("failed:%s", err)
				}
			}
		}(j)
	}
	wg.Wait()
}
