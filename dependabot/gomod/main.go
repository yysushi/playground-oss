package main

import (
	"fmt"

	"github.com/mattn/go-runewidth"
)

func main() {
	value := "ほげ"
	fmt.Printf("the length of %q is %d\n", value, runewidth.StringWidth(value))
}
