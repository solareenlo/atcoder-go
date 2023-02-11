package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	if n < a {
		fmt.Println(0)
	} else if a <= b {
		fmt.Println(n - a + 1)
	} else {
		fmt.Println((n/a-1)*b + min(n%a+1, b))
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
