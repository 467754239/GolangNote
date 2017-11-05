package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	s := Student{
		Id:   1,
		Name: "zhengyscn",
	}
	fmt.Println(s)

	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf)
	fmt.Printf("%s\n", buf)

	err = ioutil.WriteFile("/tmp/store.db", buf, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
