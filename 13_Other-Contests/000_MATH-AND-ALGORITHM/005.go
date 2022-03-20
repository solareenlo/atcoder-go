package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		sum += a
		sum %= 100
	}

	fmt.Println(sum)
}
