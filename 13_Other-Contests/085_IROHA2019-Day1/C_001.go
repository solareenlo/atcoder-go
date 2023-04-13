package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < 8; i++ {
		fmt.Println(n - 7 + i)
	}
}
