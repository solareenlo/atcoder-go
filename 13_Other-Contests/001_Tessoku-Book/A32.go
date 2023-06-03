package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	n %= a + b
	if (n-a >= 0 && n-a < a) || (n-b >= 0 && n-b < a) {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
