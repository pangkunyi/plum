package email

import (
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	err := Validate("10000@qq.com", "hello@google.com")
	if err != nil {
		t.Errorf("failed to validate email, cause by:%v", err)
	}
	err = Validate("xxxxxxxxxxxxx10000@qq.com", "hello@google.com")
	if err != nil {
		if !strings.Contains(err.Error(), "Mailbox not found") {
			t.Errorf("failed to validate email, cause by:%v", err)
		}
	}
}
