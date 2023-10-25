package main

import "fmt"

func main() {
	b := []int{4, 16, 64, 256, 1024}
	var n int
	fmt.Scan(&n)
	fmt.Println(b[n-1])
}
