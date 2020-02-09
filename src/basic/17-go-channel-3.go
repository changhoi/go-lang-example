package main

import "time"

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	ch1 <- 111
	ch1 <- 222

	ch2 <- 101
	ch2 <- 202

	close(ch1)
	close(ch2)

	/**
	수신자가 채널이 닫힐 때까지 계속 수신할 수 있는데, 방법으로는 range를 사용할 수 있다.
	*/

	// range를 사용하지 않는 경우 if로 체크할 수 있다.
	for {
		if i, succuess := <-ch1; succuess {
			println(i)
		} else {
			break
		}
	}

	// range를 사용하는 경우 훨씬 간결하게 표현할 수 있다.
	for i := range ch2 {
		println(i)
	}

	//----

	/**
	select 문은 복수 채널들을 기다리면서 준비된(데이터를 보낸) 채널을 실행하는 기능을 제공한다.
	여러 case 문에서 각각 다른 채널들을 기다리다가, 준비가 된 채널 case를 실행함
	select 문은 case 채널이 준비되지 않으면 계속 대기하고, 가장 먼저 도착한 채널의 case를 실행함

	복수 채널에서 신호가 오면, 랜덤하게 그 중 하나를 선택한다.

	select 문에 default가 존재하면 기다리지 않고 default 를 실행한다.
	*/

	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")

		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
