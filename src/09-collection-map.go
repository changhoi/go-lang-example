package main

import "fmt"

func main() {

	// map [key타입] value타입, 선언 후에는 nil map이 된다.
	var idMap map[int]string

	fmt.Println(idMap)

	if idMap == nil {
		println("Nil map")
	}

	// make로 초기화
	idMap = make(map[int]string)

	idMap[1] = "changhoi"
	idMap[2] = "dlwlrma"

	// 맵 리터럴 초기화 "map[key타입]value타입 {key: value}"
	company := map[string]string{
		"G":  "Google",
		"M":  "Microsoft",
		"GH": "Github",
	}

	fmt.Println(company)

	str := idMap[1]
	println(str)

	println("--")
	noData := idMap[3] // 값이 없는 경우에는 nil or zero 리턴 (출력은 공백을 출력함)
	println(noData)

	delete(idMap, 1) // 삭제하기
	println(idMap[1])
	println("--")

	//----

	val, exists := company["G"] // 키 값이 있는지 체크하기 위해서 읽을 때 값을 두 개 리턴한다.
	println(val, exists)

	if !exists {
		println("NO G")
	}

	//----

	myMap := map[string]string{
		"A": "apple",
		"B": "banana",
		"C": "charlie",
	}

	// Map은 Unordered이므로, 순서는 무작위
	// for range를 사용해 모든 요소를 순회할 수 있다.
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}
