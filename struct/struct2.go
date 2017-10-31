package main

import "fmt"

type Student struct {
	Id   int
	Name string
}

func main() {
	s1 := Student{
		Id:   2,
		Name: "alice",
	}

	var p *Student
	p = &s1
	p.Id = 3
	fmt.Println(s1)

	var p1 *int
	p1 = &s1.Id
	*p1 = 5
	fmt.Println(s1)

}
