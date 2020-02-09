/**
Go 루틴은 Go 런타임이 관리하는 Lightweight 논리적, 가상적 쓰레드이다. Go에서 "go" 키워드를 사용하여 함수를 호출하면, 런타임에서 새로운 goroutine을 실행한다.package src
goroutine은 비동기적으로 함수 루틴을 실행하므로 여러 코드를 동시해 실행하는데 사용된다.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	say("Sync") // 함수를 동기적으로 실행한다.

	// 함수를 비동기적으로 실행한다.
	go say("Async1")
	go say("Async2")
	go say("Async3")
	go say("Async4")

	time.Sleep(time.Second * 1) // 3초 대기하는 함수
	// 이 함수가 없으면 비동기 루틴으로 넘어간 것을 함수 끝난 것으로 인식하고 메인 함수가 종료되어버림

	//----
	var wait sync.WaitGroup
	wait.Add(1) // 1개의 고루틴을 기다리겠다는 뜻

	go func() {
		defer wait.Done()
		fmt.Println("Hello")
	}()

	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Hi")

	wait.Wait() // Go 루틴이 모두 끝날 때까지 대기

	// 1개의 고루틴만 기다리기 때문에 둘 중 하나만 Done()을 호출하는 것을 기다린다.
}
