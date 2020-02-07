package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	// 자료형 배열, 길이, 최대 용량
	slice := make([]int, 5, 10)

	// 길이, 용량
	println(len(slice), cap(slice))

	var slice2 []int

	if slice2 == nil {
		println("Nill Slice") // 길이, 용량이 0인 슬라이스 === nil
	}

	println(len(slice2), cap(slice2))
}
