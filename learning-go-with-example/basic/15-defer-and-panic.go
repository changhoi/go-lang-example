package main

import (
	"fmt"
	"os"
)

/**
defer는 특정 문장 혹은 함수를 defer를 호출하는 함수가 리턴하기 직전에 실행하게 한다.
*/

func main() {
	openFile("../readme.md")
	println("Done") // 이 문장은 패닉에 의해 실행되지 않는다 but recover가 있으면 동작하게 되어있다.

	// defer은 패닉 모드가 되어도 실행은 된다. but 콘솔 상으로 나오지 않는다... but 패닉모드가 리커버 되면 콘솔에 찍힌다.
	defer println("Main Function defer")
}

func openFile(fn string) {
	defer func() {

		// recover함수를 사용하면 panic 상태를 정상으로 되돌릴 수 있다.
		// 패닉모드에서 defer가 작동한다는 점을 이용해서 recover를 defer에 정의해두면 panic 상태를 핸들링할 수 있다
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)

	defer func() {
		fmt.Print("OpenFile Function")
		f.Close()
	}()

	if err != nil {
		panic(err)
		// 패닉 함수는 현재 실행 중인 함수를 즉시 멈추고 현재 함수의 defer를 모두 실행한 다음 즉시 리턴한다.
		// 패닉 모드의 실행 방식은 다시 상위 함수에서도 똑같이 적용되어서 계속 콜스택을 타고 올라간다.
		// 마지막에는 프로그램이 에러를 내고 종료하게 된다.
	}

}
