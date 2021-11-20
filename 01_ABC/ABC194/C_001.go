package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res, sum := 0, 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		res += a * a
		sum += a
	}

	fmt.Println(res*n - sum*sum)
}
