package strs

import (
	"testing"

	"github.com/pangkunyi/plum/strs"
)

func TestLog(t *testing.T) {
	ori := "a\n" + "\\" + "\r" + "\u0000" + "\"" + "\u0032" + "\u00a5" + "\u20a9"
	expected := "a\\n" + "\\\\" + "\\r" + "\\0" + "\\\"" + "\\Z" + "" + ""
	if expected != strs.SqlEscape(ori) {
		t.Errorf("sql escape failed[%v]-[%v]", ori, strs.SqlEscape(ori))
	}
	ori = "a\u0000" + "." + "^" + "$" + "*" + "+" + "?" + "(" + ")" + "[" + "{" + "\\" + "|"
	expected = "a\\0" + "\\." + "\\^" + "\\$" + "\\*" + "\\+" + "\\?" + "\\(" + "\\)" + "\\[" + "\\{" + "\\\\" + "\\|"
	if expected != strs.RegexEscape(ori) {
		t.Errorf("regex escape failed[%v]-[%v]", ori, strs.SqlEscape(ori))
	}
}
