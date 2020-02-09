package main

import "fmt"

func main() {

	/**
	Go에서 표준 라이브러리 패키지들은 GOROOT/pkg안에 존재한다. GOROOT 환경변수는 Go 설치 디렉토리를 가리킨다.

	일반적으로 패키지는 라이브러리로 사용되는데, main이라고 명명된 패키지는 Go Complier에 의해 특별하게 인식된다.
	컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행 프로그램으로 만든다. 그리고 이 main 패키지 안의 main 함수가 프로그램의 entry point가 된다.

	패키지를 공유 라이브러리로 만들 때에는 main 패키지나 main 함수를 사용하면 안 된다.
	*/

	fmt.Println("HELLO!")

	/**
	패키지를 사용할 때 임포트를 하는데, 임포트를 할 때 Go 컴파일러는 GOROOT 또는 GOPATH 환경 변수를 검색한다.
	표준 패키지는 GOROOT/pkg, 사용자 패키지나 3rd party 패키지는 GOPATH/pkg에서 패키지를 찾는다.

	GOPATH는 사용자가 직접 지정해줘야 한다.
	*/

	/**
	패키지 내에서 이름의 시작을 대문자로 지정하면 public으로 사용할 수 있게 되고 소문자로 시작하면 non-public으로 패키지 내부에서만 사용 가능하게 됨
	*/

}
