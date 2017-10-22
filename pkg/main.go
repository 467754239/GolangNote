package main

import (
	"fmt"

	"github.com/467754239/GolangNote/pkg/format"
)

func main() {
	format.Printline()
	s := format.PrintHtml("zhengyscn", "27")
	fmt.Println(s)
}
