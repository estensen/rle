package rle_test

import (
	"testing"

	"github.com/estensen/rle"
	"github.com/matryer/is"
)

func TestEncode(t *testing.T) {
	is := is.New(t)

	seq := "AACCCBBBBBAAAAFFFFFFFF"
	encoded := rle.Encode(seq)

	is.Equal(encoded, "A2C3B5A4F8")
}
