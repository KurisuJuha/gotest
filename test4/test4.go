package main

import "fmt"

func main()  {
	n , m := swap(10, 20)
	fmt.Println(n , m)
}

func swap(a int, b int) (int , int)  {
	return b , a
}