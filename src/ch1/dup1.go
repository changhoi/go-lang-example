package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			// fmt.Printf 함수는 포매팅된 문자열을 출력하는 함수이다.
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
