package main

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1])

	// var arr1 = [3]int{1, 2, 3}
	// var arr2 = [...]int{1, 2, 3}

	var arr3 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	println(arr3[1][2])
}
