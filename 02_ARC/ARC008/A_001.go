package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := 100 * (n / 10)

	var b int
	if n%10 >= 7 {
		b = 100
	} else {
		b = 15 * (n % 10)
	}

	fmt.Println(a + b)
}
