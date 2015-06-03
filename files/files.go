package files

import (
	"bufio"
	"os"
)

func ScanFile(filename string, lineFn func(line string) error) error {
	fd, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		if err := lineFn(scanner.Text()); err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
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
