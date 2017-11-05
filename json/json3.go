package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Student struct {
	Id   int
	Name string
}

func main() {

	content, err := ioutil.ReadFile("/tmp/store.db")
	if err != nil {
		log.Fatal(err)
	}

	var buf Student
	err = json.Unmarshal(content, &buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf)
	fmt.Printf("%#v\n", buf)
}
