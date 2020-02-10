package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// GET 호출

	/**
	http 패키지는 웹 관련 클라이언트 및 서버 기능을 제공한다.
	*/

	// 간단하게 요청을 보낼 때
	response, err := http.Get("https://github.com")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 결과 보기
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("%s\n", string(data))

	// 헤더 변경 등 비교적 세밀한 요청이 필요한 경우
	req, err := http.NewRequest("GET", "https://github.com", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "Crawler")
	client := &http.Client{}

	response2, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer response2.Body.Close()

	bytes, _ := ioutil.ReadAll(response2.Body)
	str := string(bytes)
	fmt.Println(str)
}
