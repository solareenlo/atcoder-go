package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n)
	for i := 1; i < n; i++ {
		fmt.Println(i, i+1)
	}
	fmt.Println(n, 1)
}
