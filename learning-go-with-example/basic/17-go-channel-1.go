/**
채널은 데이터를 주고받는 통로라고 할 수 있다. make 함수를 통해 미리 생성되어야 하고, 채널 연산자인 '<-, ->'를 사용해서 데이터를 주고 받는다.
흔히 고루틴 사이 데이터를 주고 받는데 사용된다. 상대편이 준비될 때까지 채널에서 대기함으로써 별도의 lock을 걸지 않고 데이터를 동기화하는데 사용된다.package src
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		ch <- 123 // 채널에 123을 보낸다.
	}()

	var i int
	println("Before <-")
	i = <-ch // 채널로부터 123을 받는다. 여기부터 ch <-123을 기다리게 된다
	println(i)

	/**
	Go 채널은 Unbuffered Channel과 Buffered Channel이 있다. 위의 예제에서 Go의 채널은 Unbuffered Channel으로, 하나의 수신자가 데이터를 받을 대까지 송시자가 데이터를 보내는 채널에 묶여있다.
	Buffer Channel을 ㅏ용하면 수신자가 받을 준비가 되어있지 않았을 지라도, 지정된 버퍼만큼 데이터를 보내고 계속 다른 일을 수행할 수 있음

	버퍼 채널은 make(chan type, N) 함수로 만들어진다. N은 사용할 버퍼의 개수
	*/

	// deadLock := make(chan int)
	// deadLock <- 1
	// fmt.Println(<-deadLock)

	// 정해둔 최대 버퍼수까지 미리 보내둘 수 있음
	ch2 := make(chan int, 1)
	ch2 <- 101 // 이건 버퍼가 되어서 fmt.Println에서 찍힘,
	// ch2 <- 102 // 이 부분은 데드락이라고 나옴
	fmt.Println(<-ch2)
	// fmt.Println(<-ch2) // 이 부분
}
