package files

import (
	"fmt"
	"log"
	"os"
	"sync/atomic"
	"time"
)

//FileLoader file load struct
type FileLoader struct {
	filename       string
	linesProcessor func([]string) (interface{}, error)
	lastModifyTime time.Time
	value          atomic.Value
}

//Value get value from file loader
func (fl *FileLoader) Value() interface{} {
	return fl.value.Load()
}

func (fl *FileLoader) start() (*FileLoader, error) {
	err := fl.loadFile()
	if err != nil {
		return fl, err
	}
	go func() {
		for {
			err := fl.loadFile()
			if err != nil {
				log.Printf("failed to load file %s\n", fl.filename)
			}
			time.Sleep(5 * time.Second)
		}
	}()
	return fl, nil
}

func (fl *FileLoader) loadFile() error {
	fi, err := os.Stat(fl.filename)
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return fmt.Errorf("failed to load file, cause by [%s] is dir", fl.filename)
	}
	if fi.ModTime().After(fl.lastModifyTime) {
		lines := make([]string, 0)
		if err = ScanFile(fl.filename, func(line string) error {
			lines = append(lines, line)
			return nil
		}); err != nil {
			return err
		}
		val, err := fl.linesProcessor(lines)
		if err != nil {
			return err
		}
		fl.value.Store(val)
		fl.lastModifyTime = fi.ModTime()
	}
	return nil
}

//NewFileLoader create a file loader instance
func NewFileLoader(filename string, linesProcessor func([]string) (interface{}, error)) (*FileLoader, error) {
	fileLoader := &FileLoader{filename: filename, linesProcessor: linesProcessor}
	return fileLoader.start()
}
