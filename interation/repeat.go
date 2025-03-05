package interation

import "strings"

func Repeat(character string, count int) string {
	var repeated strings.Builder
	for range count {
		repeated.WriteString(character)
	}

	return repeated.String()
}
