package main

import "fmt"

// struct를 정의
type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

// 포인터를 리턴한다는 표시
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d // 포인터를 전달하는 방법
}

func main() {
	// person 객체 생성, 필드값 설정
	p := person{}
	p.name = "Lee"
	p.age = 10
	fmt.Println(p)

	var p1 person
	p1 = person{"Gildong", 20}
	p2 := person{"Changhoi", 25}

	// 위와 같은 방식은 순서대로 할당이 되고, 아래와 같은 경우는 순서와 상관 없이 넣을 수 있다.
	// 만약 필드가 생략되면 Zero value (0, 0.0, "", nil, ...)를 갖는다
	p3 := person{name: "Sean", age: 50}

	fmt.Println(p1, p2, p3)

	// new 함수를 사용하면 모든 필드를 Zero Valu로 초기화하고 person 객체의 포인터 *person을 리턴한다. 객체 포인터에서 필드 엑세스를 할 때 '.'을 사용하는데 포인터가 자동으로 Dereference 된다.
	p4 := new(person)
	p4.name = "Lee" // 여기서 p4는 포인터지만 접근할 때는 마찬가지로 '.'을 사용한다.

	fmt.Println(p4)
	fmt.Println(&p3)

	/**
	Go에서 struct를 변경할 때 새로운 객체가 생성되는 것이 아니라 내부에서 값의 변화만 이루어진다. (일반적으로 JS, Python과 마찬가지)
	하지만 파라미터로 전달 되는 경우에는 struct가 복사되어 새로운 객체가 넘어간다. (Call By Value) 따라서 레퍼런스를 전달하려고 할 때는 struct의 포인터를 전달해주어야 한다.
	new를 쓰거나, &p1 ?
	*/

	//----

	/**
	구조체 필드가 사용 전에 초기화 되어야 하는 경우 생성자 함수를 사용할 수 있다.
	생성자 함수는 struct를 리턴하는 함수이다.
	*/

	dic := newDict() // 생성자 호출
	dic.data[1] = "A"

	fmt.Println(dic, *dic, dic.data)
}
