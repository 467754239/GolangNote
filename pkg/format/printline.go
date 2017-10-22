package format

import "fmt"

func Printline() {
	fmt.Println("func Print")
}

func PrintHtml(name string, age string) string {
	return fmt.Sprintf("<h1>name:%s, age:%s</h1>", name, age)
}
