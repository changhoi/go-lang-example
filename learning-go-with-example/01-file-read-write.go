package main

import (
	"io"
	"io/ioutil"
	"os"
)

const BASE_DIR = "/Users/changhoi/dev/changhoi/go-dev/"

func main() {
	/**
	os.Open() 함수는 기존 파일을 열 때 사용.
	os.Create() 함수는 새 파일을 생성할 때 사용한다.
	*/
	fi, err := os.Open(BASE_DIR + "README.md")
	// 파일 열기

	if err != nil {
		panic(err)
	}
	defer fi.Close()

	// 출력 파일 생성
	fo, err := os.Create(BASE_DIR + "README_COPY.md")

	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 1024)

	for {
		cnt, err := fi.Read(buff) // 읽기 (버퍼에 읽어옴)
		if err != nil && err != io.EOF {
			panic(err)
		}

		// 끝이면 루프를 종료
		if cnt == 0 {
			break
		}

		// 쓰기
		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}

	/**
	ioutil 패키ㅣ는 i/o 관련한 편리한 유틸리티를 제공하는 패키지이다
	입력 파일이 매우 크지 않은 경우, 이 패키지의 ReadFile, WriteFile 함수를 이용해서 편리하게 파일을 읽고 쓸 수 있다.
	*/

	// 파일 읽기
	file, err := ioutil.ReadFile(BASE_DIR + "README.md")

	if err != nil {
		panic(err)
	}

	// 파일 쓰기, 뒤에 0644는 file permission임. 윈도우의 경우에 달라짐
	err = ioutil.WriteFile(BASE_DIR+"README_COPY_2.md", file, 0644)

	if err != nil {
		panic(err)
	}
}
