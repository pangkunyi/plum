package files

import (
	"fmt"
	"log"
	"os"
	"sync/atomic"
	"time"
)

type FileLoader struct {
	filename       string
	linesProcessor func([]string) (interface{}, error)
	lastModifyTime time.Time
	value          atomic.Value
}

func (this *FileLoader) Value() interface{} {
	return this.value.Load()
}

func (this *FileLoader) start() (*FileLoader, error) {
	err := this.loadFile()
	if err != nil {
		return this, err
	}
	go func() {
		for {
			err := this.loadFile()
			if err != nil {
				log.Printf("failed to load file %s\n", this.filename)
			}
			time.Sleep(5 * time.Second)
		}
	}()
	return this, nil
}
func (this *FileLoader) loadFile() error {
	fi, err := os.Stat(this.filename)
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return fmt.Errorf("failed to load file, cause by [%s] is dir", this.filename)
	}
	if fi.ModTime().After(this.lastModifyTime) {
		lines := make([]string, 0)
		if err = ScanFile(this.filename, func(line string) error {
			lines = append(lines, line)
			return nil
		}); err != nil {
			return err
		}
		val, err := this.linesProcessor(lines)
		if err != nil {
			return err
		}
		this.value.Store(val)
		this.lastModifyTime = fi.ModTime()
	}
	return nil
}

func NewFileLoader(filename string, linesProcessor func([]string) (interface{}, error)) (*FileLoader, error) {
	fileLoader := &FileLoader{filename: filename, linesProcessor: linesProcessor}
	return fileLoader.start()
}
