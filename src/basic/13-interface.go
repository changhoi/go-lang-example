package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// 인터페이스를 struct가 구현하도록 강제하는 기능은 하지 않는 것 같다. 다만 저런 인터페이스를 갖춰야만 해당 인터페이스를 인자값으로 받는 경우에 통과된다.

func main() {
	r := Rect{10., 20.}
	c := Circle{10.}

	showArea(r, c)

	var x interface{}
	x = 1
	x = "Changhoi"
	printIt(x) // interface{}는 어떤 값이든 올 수 있다는 것을 의미한다 마지막 값을 출력하고 있다.

	var a interface{} = 1
	i := a
	j := a.(int) // a.(T) a는 nil이 아니고, x는 T 타입에 속한다는 점을 확언 하는 표현 (Type Assertion)

	println(i)
	println(j)
}

func printIt(v interface{}) {
	fmt.Println(v)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		println(a)
	}
}
