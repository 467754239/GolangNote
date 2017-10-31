package main

import "fmt"

type Student struct {
	Id   int
	Name string
}

func main() {
	// struct 初始化
	var s Student
	s.Id = 1
	s.Name = "jack"
	fmt.Println(s)

	s1 := Student{
		Id:   2,
		Name: "jack",
	}
	fmt.Printf("%#v\n", s1)

	s1 = s
	fmt.Printf("%#v\n", s1)
}
