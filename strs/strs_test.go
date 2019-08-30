package strs

import (
	"testing"
)

func TestSQLEscape(t *testing.T) {
	ori := "a\n" + "'" + "\\" + "\r" + "\u0000" + "\"" + "\u0032" + "\u00a5" + "\u20a9"
	expected := "a\\n" + "\\'" + "\\\\" + "\\r" + "\\0" + "\\\"" + "\\Z" + "" + ""
	if expected != SQLEscape(ori) {
		t.Errorf("sql escape failed[%v]-[%v]", ori, SQLEscape(ori))
	}
}

func TestRegexEscape(t *testing.T) {
	ori := "a\u0000" + "." + "^" + "$" + "*" + "+" + "?" + "(" + ")" + "[" + "{" + "\\" + "|"
	expected := "a\\0" + "\\." + "\\^" + "\\$" + "\\*" + "\\+" + "\\?" + "\\(" + "\\)" + "\\[" + "\\{" + "\\\\" + "\\|"
	if expected != RegexEscape(ori) {
		t.Errorf("regex escape failed[%v]-[%v]", ori, RegexEscape(ori))
	}
}
