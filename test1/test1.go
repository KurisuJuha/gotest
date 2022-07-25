package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Println(omikuzi(rand.Intn(6)))
	}
}

func omikuzi(a int) string {
	var s string

	switch a + 1 {
	case 6:
		s = "大吉"
	case 5:
		fallthrough
	case 4:
		s = "中吉"
	case 3, 2:
		s = "吉"
	case 1:
		s = "凶"
	}

	return s
}
