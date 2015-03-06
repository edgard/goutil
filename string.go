package goutil

import "strings"

// StringCap limits string to number or characters
func StringCap(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return strings.TrimSpace(s[0:length-3]) + "..."
}
