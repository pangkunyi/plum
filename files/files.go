package files

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

//ScanFileFull scan a file line by line, it can handle the long line
func ScanFileFull(filename string, lineFn func(line string) error) error {
	fd, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)
	lineBuf := bytes.NewBuffer([]byte{})
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
			if lineBuf.Len() > 0 {
				lineBuf.Write(line)
				if err := lineFn(lineBuf.String()); err != nil {
					return err
				}
				lineBuf.Reset()
			} else {
				if err := lineFn(string(line)); err != nil {
					return err
				}
			}
		} else {
			lineBuf.Write(line)
		}
	}
	return nil
}

//ScanFilesFull scan files line by line, it can handle the long line
func ScanFilesFull(filenames []string, lineFn func(line string) error) error {
	for _, filename := range filenames {
		if err := ScanFileFull(filename, lineFn); err != nil {
			return err
		}
	}
	return nil
}

//ScanFile scan a file line by line, it can not handle the long line
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

//ScanFiles scan files line by line, it can not handle the long line
func ScanFiles(filenames []string, lineFn func(line string) error) error {
	for _, filename := range filenames {
		if err := ScanFile(filename, lineFn); err != nil {
			return err
		}
	}
	return nil
}
