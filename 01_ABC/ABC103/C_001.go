package main

import "fmt"

func main() {
	var n, a, sum int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		sum += a - 1
	}
	fmt.Println(sum)
}
