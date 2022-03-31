package main

import "fmt"

func f(x int) int {
	return x*x + 2*x + 3
}

func main() {
	var t int
	fmt.Scan(&t)

	fmt.Println(f(f(f(t)+t) + f(f(t))))
}
