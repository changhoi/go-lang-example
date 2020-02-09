/**
일반적인 OOP를 지원하는 언어들과 다르게 Go에서는 struct가 필드만 가지고, 메서드는 별도로 분리되어 작성된다.
Go에서 Method라는 것은 특별한 형태의 Func 함수이다. 함수 정의에서 func 키워드와 함수명 사이에 해당 함수가 어떤 struct를 위한 메서드인지 표기하는 것이다.package src
이때 어떤 struct를 위한 것인지를 표기하는 부분을 receiver라고 부른다

receiver는 funct `(struct변수명 struct타입) 메서드이름 리턴타입 {}`와 같이 작성되고 이 때 변수명 부분은 함수 내에서 마치 인자값처럼 사용된다.
*/

package main

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	r.width++
	return r.width * r.height
}

func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()
	println(area)

	rect2 := Rect{10, 20}
	area2 := rect2.area2()
	println(rect.width, area)
	println(rect2.width, area2)
}

/**
Value vs 포인터 receiver

Value Receiver는 struct의 데이터를 복사해서 전달하고, Pointer Receiver는 포인터만을 전달함
즉 메서드를 호출 했을 때, JS의 this에 해당하는 부분이 어떻게 전달 되는지를 의미한다.

Value Receiver의 경우에는 메서드 내부에서 아무리 값을 변경 해주더라도 this의 값은 변하지 않는다.
Pointer Receiver의 경우에는 메서드 내부에서 struct의 값을 변경하면 this의 값이 바뀌게 된다.
*/
