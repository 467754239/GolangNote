package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Id   int
	Name string
}

func main() {

	str := `{"id":1,"name":"zhengyscn"}`

	var buf Student
	err := json.Unmarshal([]byte(str), &buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf)
	fmt.Printf("%#v\n", buf)
}
