package yeet

import (
	_ "embed"
	"fmt"
)

//go:embed richard.txt
var richard string

func Do() {
	fmt.Println(richard)
}
