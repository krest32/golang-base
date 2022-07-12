package main

import "fmt"

// if....else
func main() {
	var a int = 100
	fmt.Printf("a = %d\n", a)

	if a < 20 {
		fmt.Printf("a < 20\n")
	} else {
		fmt.Printf("a >= 20\n")
	}

}
