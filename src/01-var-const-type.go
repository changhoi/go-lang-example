package main

func main() {
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(u)

	println(i, u, f)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)

	println(str, bytes, str2)
}
