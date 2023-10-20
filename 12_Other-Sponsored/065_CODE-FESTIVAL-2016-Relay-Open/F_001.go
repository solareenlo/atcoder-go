package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(f(x))
}

func f(x int) int {
	if x == 0 {
		return 2
	} else {
		return f(x-1)*2 + 2
	}
}
