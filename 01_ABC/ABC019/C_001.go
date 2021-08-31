package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	m := map[int]struct{}{}
	var a int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		for a%2 == 0 {
			a /= 2
		}
		m[a] = struct{}{}
	}
	fmt.Println(len(m))
}
