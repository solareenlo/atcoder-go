package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	m := map[int]struct{}{}
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		m[tmp] = struct{}{}
	}
	fmt.Println(len(m))
}
