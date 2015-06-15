package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	ERR_CMD_FAILURE = "程序[%v %v]执行失败, 原因：\n%v"
	ERR_CMD_TIMEOUT = "程序[%v %v]执行超时, 原因：\n%v"
)

func Cmd(name string, arg ...string) (string, error) {
	return CmdW("", name, arg...)
}

func CmdW(workDir string, name string, arg ...string) (string, error) {
	log.Printf("执行程序[%v %v]...\n", name, arg)
	cmd := exec.Command(name, arg...)
	if workDir != "" {
		if fi, err := os.Stat(workDir); err != nil || !fi.IsDir() {
			return "", fmt.Errorf(ERR_CMD_FAILURE, name, arg, workDir+"目录不存在")
		}
		cmd.Dir = workDir
	}
	out, err, timeout := CmdRunWithTimeout(cmd, 10*60*time.Second)
	if timeout {
		log.Printf("超时\n")
		return string(out), fmt.Errorf(ERR_CMD_FAILURE, name, arg, err.Error())
	} else if err != nil {
		log.Printf("失败\n")
		return string(out), fmt.Errorf(ERR_CMD_FAILURE, name, arg, err.Error())
	} else {
		log.Printf("成功\n")
		return string(out), nil
	}
}

func CmdRunWithTimeout(cmd *exec.Cmd, timeout time.Duration) ([]byte, error, bool) {
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	if err := cmd.Start(); err != nil {
		return b.Bytes(), err, false
	}
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		// timeout
		if err = cmd.Process.Kill(); err != nil {
			log.Printf("failed to kill: %s, error: %s", cmd.Path, err)
		}
		go func() {
			<-done // allow goroutine to exit
		}()
		log.Printf("process:%s killed", cmd.Path)
		return b.Bytes(), err, true
	case err = <-done:
		return b.Bytes(), err, false
	}
}
