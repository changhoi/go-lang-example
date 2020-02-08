package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)

	//----

	ch2 := make(chan int, 2)

	ch2 <- 111
	ch2 <- 222

	/**
	채널을 닫으면 송신을 할 수는 없게 된다. 하지만 수신은 가능함. <-ch 에서는 두 개의 리턴 값을 갖는데 하나는 채널의 메시지, 두 번째는 수신이 잘 되었는지 나타내는 boolean 값이다.
	채널이 닫혀 있으면 두 번째 리턴 값은 false가 된다. 있는 경우엔 true (더 이상 수신할 message가 없다면, 첫 번째 값은 zero value가 리턴된다)

	*/

	close(ch2) //채널을 닫는다.

	println(<-ch2)
	println(<-ch2) //채널 수신

	if message, success := <-ch2; !success {
		fmt.Println("message:", message, "success:", success)
		println("더 이상 데이터 없음")
	}
}

/**
채널을 함수 파라미터로 전달할 때 보통은 송, 수신을 모두 하는 채널을 전달하지만, 특별이 인자로 받은 채널로 송신만 할 것인지 수신만 할 것인지를 지정할 수도 있음
*/

// 파라미터로 받은 채널이 송신하는 경우 파라미터
func sendChan(ch chan<- string) {
	ch <- "Data"
}

// 파라미터로 받은 채널이 수신만 하는 경우
func receiveChan(ch <-chan string) {
	// ch <- "Data" ** 오류가 남
	data := <-ch
	fmt.Println(data)
}
