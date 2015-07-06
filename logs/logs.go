package logs

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	fd       *os.File
	mu       sync.Mutex
	filename string
	rotate   bool
	logDay   int
}

func NewLogger(filename string, rotate bool) *Logger {
	return &Logger{filename: filename, rotate: rotate, logDay: -1}
}

func (l *Logger) checkLogFile() error {
	now := time.Now()
	if l.fd != nil {
		if !l.rotate || l.logDay == now.Day() {
			return nil
		}
		l.fd.Close()
	}
	filename := l.filename
	if l.rotate {
		filename = fmt.Sprintf(l.filename, now.Format("2006-01-02"))
	}
	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	l.fd = fd
	return err
}

func (l *Logger) Print(msg string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if err := l.checkLogFile(); err != nil {
		return err
	}
	_, err := l.fd.Write([]byte(msg))
	return err
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.Print(fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(msg string) {
	l.Print(msg)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Print(fmt.Sprintf(format, v...))
	os.Exit(1)
}
