package main

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	member := Member{"Alex", 10, true}

	// JSON 마샬링이라고 부른다 (인코딩)
	jsonBytes, err := json.Marshal(member)
	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	// 테스트용 데이터
	jsonBytes2, _ := json.Marshal(Member{"Tim", 1, true})

	var mem Member

	// JSON 디코딩 하는 부분, 언마샬링
	err2 := json.Unmarshal(jsonBytes2, &mem)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(mem.Name, mem.Age, mem.Active)
}
