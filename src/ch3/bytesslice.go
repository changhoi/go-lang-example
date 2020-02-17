package main

import "fmt"

func main() {
	var s []byte
	s = []byte("hello")
	fmt.Println(s)

	s[0] = []byte("H")[0]
	fmt.Println(s)
}
