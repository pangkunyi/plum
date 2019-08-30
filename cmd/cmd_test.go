package cmd

import (
	"testing"
)

func TestCmd(t *testing.T) {
	_, err := Execute("ls")
	if err != nil {
		t.Errorf("failed:%s", err)
	}
}
