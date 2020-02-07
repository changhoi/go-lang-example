package main

func main() {
	a := 1 + 5
	b := 5 - 3
	c := 10 / 3
	d := 10 / 2
	e := 10 % 5
	f := 10 % 3
	g := 10 % 3

	println(a, b, c, d, e, f)

	println(f == g, f != e, e == g)
	println(f < e, f >= g)

	pointer := &a
	println(pointer, *pointer)

	a = 11
	print(pointer, *pointer)
}
