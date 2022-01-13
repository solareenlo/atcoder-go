package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	b := make([]int, n)
	x := 0
	for i := range b {
		fmt.Scan(&b[i])
		x ^= b[i]
	}

	if x != 0 {
		fmt.Println(-1)
		return
	}

	res := 0
	for i := 0; i < n; i++ {
		fmt.Println(res)
		res ^= b[i]
	}
}
