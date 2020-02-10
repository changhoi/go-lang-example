package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	reqBody := bytes.NewBufferString("Post plain text")
	// 간단한 데이터를 post할 때 사용
	// URL, MIME 타입, BODY
	resp, err := http.Post("http://httpbin.org/post", "text/plain", reqBody)

	if err != nil {
		panic(err)
	}

	respBody, err1 := ioutil.ReadAll(resp.Body)
	if err1 == nil {
		str := string(respBody)
		println(str)
	}

	/**
	http.PostForm()은 폼데이터를 보낼 때 유용한 메서드이다. 일반 페이지에서 Submit 할 때처럼 Form 데이터를 만들어서 서버로 전송한다.
	*/

	resp2, err2 := http.PostForm("http://httpbin.org/post", url.Values{"Name": {"Lee"}, "Age": {"10"}})
	if err2 != nil {
		panic(err2)
	}
	defer resp2.Body.Close()

	responseBody, err3 := ioutil.ReadAll(resp2.Body)
	if err3 == nil {
		str := string(responseBody)
		println(str)
	}

	/**
	JSON 데이터 POST하기는 위와 비슷하지만, MIME 타입에 application/json을 적어두고, JSON으로 인코딩된 데이터를 바디로 전달해주면 된다.
	*/

	person := Person{"Alex", 10}
	personBytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(personBytes)

	resp3, err4 := http.Post("http://httpbin.org/post", "application/json", buff)
	responseBody2, err4 := ioutil.ReadAll(resp3.Body)
	if err4 != nil {
		str := string(responseBody2)
		println(str)
	}

	/**
	마찬가지로 세밀한 제어를 하기 위해서 request 객체를 만들 수도 있다.
	*/

	req, err5 := http.NewRequest("POST", "http://httpbin.org/post", buff)
	if err5 != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp3, err6 := client.Do(req)
	if err6 != nil {
		panic(err)
	}
	defer resp3.Body.Close()

	responseBody3, err7 := ioutil.ReadAll(resp3.Body)
	if err7 == nil {
		str := string(responseBody3)
		println(str)
	}
}
