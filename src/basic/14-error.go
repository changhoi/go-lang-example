package main

import (
	"log"
	"os"
)

func main() {
	/**
	os.Open 함수는 리턴값이 file *File과 err error 두 가지인데, 이때 err error가 존재하면 nil이 아니기 때문에 nil을 에러가 있는 건지 찾아볼 수 있다.
	*/
	f, err := os.Open("/Users/changhoi/dev/changhoi/go-dev/README.md")
	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())

	// 에러를 만들 때는 내장된 인터페이스를 구현하는 것으로 만들어지기 때문에 아래와 같이 사용자가 정의한 에러를 받거나 모든 에러를 받는 가장 상위 에러 타입을 해둘 수도 있음

	switch err.(type) {
	default: // no error
	case MyErrror:
		log.Print("Log My Error")
	case error:
		log.Fatal(err.Error())
	}
}
