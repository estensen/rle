package rle_test

import (
	"testing"

	"github.com/estensen/rle"
	"github.com/matryer/is"
)

func TestEncode(t *testing.T) {
	is := is.New(t)

	tests := []struct {
		name     string
		seq      string
		expected string
	}{
		{
			name:     "emtpy",
			seq:      "",
			expected: "",
		},
		{
			name:     "basic",
			seq:      "AACCCBBBBBAAAAFFFFFFFF",
			expected: "A2C3B5A4F8",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			encoded := rle.Encode(tc.seq)
			is.Equal(encoded, tc.expected)
		})
	}
}

func TestDecode(t *testing.T) {
	is := is.New(t)

	tests := []struct {
		name     string
		encoded  string
		expected string
	}{
		{
			name:     "empty",
			encoded:  "",
			expected: "",
		},
		{
			name:     "basic",
			encoded:  "A2C3B5A4F8",
			expected: "AACCCBBBBBAAAAFFFFFFFF",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			decoded := rle.Decode(tc.encoded)
			is.Equal(decoded, tc.expected)
		})
	}
}
