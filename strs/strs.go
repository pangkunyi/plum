package strs

import "strings"

var sqlReplacer = strings.NewReplacer("\n", "\\n", "'", "\\'", "\\", "\\\\", "\r", "\\r", "\u0000", "\\0", "\"", "\\\"", "\u0032", "\\Z", "\u00a5", "", "\u20a9", "")

var regexReplacer = strings.NewReplacer("\u0000", "\\0", ".", "\\.", "^", "\\^", "$", "\\$", "*", "\\*", "+", "\\+", "?", "\\?", "(", "\\(", ")", "\\)", "[", "\\[", "{", "\\{", "\\", "\\\\", "|", "\\|")

func SqlEscape(s string) string {
	return sqlReplacer.Replace(s)
}

func RegexEscape(s string) string {
	return regexReplacer.Replace(s)
}
