package rle_test

import (
	"testing"

	"github.com/estensen/rle"
	"github.com/matryer/is"
)

func TestEncode(t *testing.T) {
	is := is.New(t)

	seq := "AACCCBBBBBAAAAFFFFFFFF"
	expected := "A2C3B5A4F8"
	encoded := rle.Encode(seq)

	is.Equal(encoded, expected)
}

func TestDecode(t *testing.T) {
	is := is.New(t)

	encoded := "A2C3B5A4F8"
	expected := "AACCCBBBBBAAAAFFFFFFFF"
	decoded := rle.Decode(encoded)

	is.Equal(decoded, expected)
}
