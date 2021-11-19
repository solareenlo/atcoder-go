package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	m := map[int]struct{}{}
	for i := 2; i*i < n+1; i++ {
		for j := i * i; j < n+1; j *= i {
			m[j] = struct{}{}
		}
	}

	fmt.Println(n - len(m))
}
