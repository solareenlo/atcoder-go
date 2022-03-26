package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := 0
	for n = n + 1; n > 9; n /= 10 {
		a += 9
	}
	fmt.Println(a + n - 1)
}
