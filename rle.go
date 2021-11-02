package rle

import (
	"fmt"
	"strings"
)

func Encode(seq string) string {
	var compressed strings.Builder
	count := 1
	char := seq[0]

	for i := 1; i < len(seq); i++ {
		fmt.Println(i)
		if seq[i] == char {
			count += 1
		} else {
			compressed.WriteString(fmt.Sprintf("%c%d", char, count))
			char = seq[i]
			count = 1
		}
	}

	compressed.WriteString(fmt.Sprintf("%c%d", char, count))
	return compressed.String()
}
