package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	m := []int{6, 2, 4, 8}
	fmt.Println(m[n%4])
}
