package main

func main() {
	total, sum := sum(1, 7, 3, 5, 9)
	println(total, sum)
}

func sum(nums ...int) (total int, count int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
