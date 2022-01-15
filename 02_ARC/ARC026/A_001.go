package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	res := 0
	if n <= 5 {
		res += n * b
	} else {
		res += 5 * b
		res += (n - 5) * a
	}

	fmt.Println(res)
}
