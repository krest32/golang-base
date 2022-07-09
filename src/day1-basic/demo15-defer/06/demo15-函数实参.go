package main

// 函数实参
func test15() (int, int) {
	return 1, 2
}

func add15(x, y int) int {
	return x + y
}

func sum15(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}

	return x
}

func main() {
	println(add15(test15()))
	println(sum15(test15()))
}