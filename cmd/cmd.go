package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	errCmdFailure = "command[%v %v] execute failed, cause by:\n%v"
	errCmdTimeout = "command[%v %v] execute timeout, cause by:\n%v"
)

//Execute execute command
func Execute(name string, arg ...string) (string, error) {
	return ExecuteW("", name, arg...)
}

//ExecuteW execute command in specified work directory
func ExecuteW(workDir string, name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	if workDir != "" {
		if fi, err := os.Stat(workDir); err != nil || !fi.IsDir() {
			return "", fmt.Errorf(errCmdFailure, name, arg, workDir+" not exists!")
		}
		cmd.Dir = workDir
	}
	out, timeout, err := ExecuteWithTimeout(cmd, 10*60*time.Second)
	if timeout {
		return string(out), fmt.Errorf(errCmdFailure, name, arg, err.Error())
	} else if err != nil {
		return string(out), fmt.Errorf(errCmdFailure, name, arg, err.Error())
	} else {
		return string(out), nil
	}
}

//ExecuteWithTimeout execute command with timeout
func ExecuteWithTimeout(cmd *exec.Cmd, timeout time.Duration) ([]byte, bool, error) {
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	if err := cmd.Start(); err != nil {
		return b.Bytes(), false, err
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
			err = fmt.Errorf("failed to kill: %s, cause by: %s", cmd.Path, err)
		}
		go func() {
			<-done // allow goroutine to exit
		}()
		return b.Bytes(), true, err
	case err = <-done:
		return b.Bytes(), false, err
	}
}
