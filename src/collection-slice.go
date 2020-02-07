package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	// 자료형 배열, 길이, 최대 용량
	slice := make([]int, 5, 10)
}
