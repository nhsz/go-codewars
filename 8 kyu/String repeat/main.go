package kata

import "strings"

// RepeatStr https://www.codewars.com/kata/57a0e5c372292dd76d000d7e
func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}
