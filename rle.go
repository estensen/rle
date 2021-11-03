package rle

import (
	"fmt"
	"strconv"
	"strings"
)

func Encode(seq string) string {
	if seq == "" {
		return ""
	}

	var encoded strings.Builder
	count := 1
	char := seq[0]

	for i := 1; i < len(seq); i++ {
		if seq[i] == char {
			count += 1
		} else {
			encoded.WriteString(fmt.Sprintf("%c%d", char, count))
			char = seq[i]
			count = 1
		}
	}

	encoded.WriteString(fmt.Sprintf("%c%d", char, count))
	return encoded.String()
}

func Decode(encoded string) string {
	if encoded == "" {
		return ""
	}

	var decoded strings.Builder

	for i := 1; i <= len(encoded); i += 2 {
		char := encoded[i-1]
		count, _ := strconv.Atoi(string(encoded[i]))
		for j := 0; j < count; j++ {
			decoded.WriteString(fmt.Sprintf("%c", char))
		}
	}

	return decoded.String()
}
