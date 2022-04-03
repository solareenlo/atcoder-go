package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	m := make(map[int]int)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		m[x]++
	}

	fmt.Println(len(m))
}
