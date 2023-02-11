package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 1
	for i := 1; i < n+1; i++ {
		res *= i
	}
	fmt.Println(res)
}
