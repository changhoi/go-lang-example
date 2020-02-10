package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	// range는 값의 쌍 (인덱스와 그 값)을 리턴하기 때문에 사용하지 않는 인덱스는 '_'로 처리 해서 에러가 나는 것을 막을 수 있다.
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
