package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x := 0
	for n > 0 {
		x += (n % 10)
		n /= 10
	}

	if x == 1 {
		x *= 10
	}
	fmt.Println(x)
}
