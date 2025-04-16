package utils

import "strings"

func SplitArgs(s string) []string {
	return strings.Fields(s)
}
