package main

import "fmt"

func main() {
	var k, a, b int
	fmt.Scan(&k, &a, &b)
	if a >= k {
		fmt.Println(1)
		return
	}
	if a <= b {
		fmt.Println(-1)
		return
	}
	n := 0
	if (k-a)%(a-b) >= 1 {
		n = 1
	}
	fmt.Println(((k-a)/(a-b)+n)*2 + 1)
}
