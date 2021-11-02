package main

import (
	"fmt"

	"github.com/estensen/rle"
)

func main() {
	fmt.Println(rle.Encode("AACCCBBBBBAAAAFFFFFFFF"))
}
