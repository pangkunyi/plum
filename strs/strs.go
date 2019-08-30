package strs

import "strings"

var sqlReplacer = strings.NewReplacer("\n", "\\n", "'", "\\'", "\\", "\\\\", "\r", "\\r", "\u0000", "\\0", "\"", "\\\"", "\u0032", "\\Z", "\u00a5", "", "\u20a9", "")

var regexReplacer = strings.NewReplacer("\u0000", "\\0", ".", "\\.", "^", "\\^", "$", "\\$", "*", "\\*", "+", "\\+", "?", "\\?", "(", "\\(", ")", "\\)", "[", "\\[", "{", "\\{", "\\", "\\\\", "|", "\\|")

//SQLEscape escape the sql string to a validated sql string
func SQLEscape(sql string) string {
	return sqlReplacer.Replace(sql)
}

//RegexEscape escape the regex string to a validate regex string
func RegexEscape(re string) string {
	return regexReplacer.Replace(re)
}
