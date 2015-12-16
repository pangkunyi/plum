package files

import (
	"bufio"
	"io"
	"os"
)

func ScanFile(filename string, lineFn func(line string) error) error {
	fd, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		if !isPrefix {
			if err := lineFn(string(line)); err != nil {
				return err
			}
		}
	}
	return nil
}

func ScanFiles(filenames []string, lineFn func(line string) error) error {
	for _, filename := range filenames {
		if err := ScanFile(filename, lineFn); err != nil {
			return err
		}
	}
	return nil
}
