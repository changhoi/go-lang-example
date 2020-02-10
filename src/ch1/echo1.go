package main

import (
	"fmt"
	"os"
)

// 커멘드라인 인수를 출력한다.
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
