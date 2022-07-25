package main

import "fmt"

type Student struct {
	Name  string
	Class int
	No    int
}

func main() {
	var p Student
	p.Name = "Tom"
	p.Class = 1
	p.No = 1

	q := struct {
		name string
		age  int
	}{
		name: "Gopher",
		age:  10,
	}

	q.age++

	fmt.Println(p)
	fmt.Println(q)
}
