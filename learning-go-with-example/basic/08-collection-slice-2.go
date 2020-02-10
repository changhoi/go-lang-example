package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	s = s[2:5]
	fmt.Println(s)

	s2 := []int{0, 1, 2, 3, 4, 5}

	s2 = s2[2:5]
	fmt.Println(s2)

	s2 = s2[1:]
	fmt.Println(s2)

	s2 = s2[:]
	fmt.Println(s2)

	s2 = s2[:1]
	fmt.Println(s2)

	//----

	slice := []int{0, 1}

	slice = append(s, 2)
	fmt.Println(slice)

	slice = append(s, 3, 4, 5)
	fmt.Println(slice)

	/**
	용량이 남아있다면 용량 내에서 length를 변경해서 데이터를 추가.

	용량을 초과하는 경우 현재 용량의 2배에 해당하는 새로운 Underlying array를 생성하고,
	기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당.
	*/

	sliceA := make([]int, 0, 3)
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA)

	//----

	sliceB := []int{1, 2, 3}
	sliceC := []int{4, 5, 6}

	sliceD := append(sliceB, sliceC...)
	fmt.Println(sliceD)

	//----

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)
	println(len(target), cap(target))

	/**
	슬라이스는 실제 배열을 가리키는 포인터 정보만 가진다.
	따라서, 여기서의 복제는 소스 슬라이스가 갖는 배열의 데이터를 타겟 슬라이스가 갖는 배열로 복제한다.
	*/

	/**
	슬라이스는 내부적으로 사용하는 배열의 부분영역인 세그먼트에 대한 메타 정보를 가지고 있다.
	슬라이스는 크게 3가지 필드로 구성되어 있음:
		1) 배열에 대한 포인터 정보
		2) 세그먼트의 길이
		3) 세그먼트의 최대 용량

	처음 슬라이스가 생성될 때
		1) 길이와 용량이 지정 된 경우:
			- 용량만큼 배열을 생성
			- 슬라이스 첫 번째 필드에 해당 배열의 처음 메모리 위치를 지정
			- 두 번째 길이 필드는 지정된 길이
			- 세 번째 용량 필드는 전체 배열의 크기를 갖게 됨

	@example: http://golang.site/images/basics/go-slice-internal.png
	*/
}
