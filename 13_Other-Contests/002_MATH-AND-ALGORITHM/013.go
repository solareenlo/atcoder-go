package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	m := divList(n)

	for k, _ := range m {
		fmt.Println(k)
	}
}

func divList(n int) map[int]struct{} {
	div := map[int]struct{}{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			div[i] = struct{}{}
			if i*i != n {
				div[n/i] = struct{}{}
			}
		}
	}
	return div
}
