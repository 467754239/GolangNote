package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Wc(s []string) map[string]int {
	m := make(map[string]int)
	for _, v := range s {
		_, ok := m[v]
		if ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}
	return m

}

func main() {
	// 词频统计
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("File contents: %s", content)
	data := strings.Fields(string(content))
	m := Wc(data)

	for word, cnt := range m {
		fmt.Println(word, cnt)
	}
}
