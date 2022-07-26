package main

import "os"

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

func main() {
	sf, err := os.Open("test5.go")
	print(sf, err)
	defer sf.Close()
}
