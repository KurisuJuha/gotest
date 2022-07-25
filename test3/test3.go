package main

import "fmt"

func main() {
	var ns = [5]int{10, 20, 30, 40, 50}
	fmt.Println(ns)

	ns2 := [...]int{5: 50, 100: 100}
	fmt.Println(ns2)

	println(ns2[3])
	println(len(ns))
	fmt.Println(ns[1:2])
}
