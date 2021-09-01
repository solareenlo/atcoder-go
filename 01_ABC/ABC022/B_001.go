package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n)

	m := map[int]bool{}
	res := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if m[a] {
			res++
		}
		m[a] = true
	}
	fmt.Println(res)
}
