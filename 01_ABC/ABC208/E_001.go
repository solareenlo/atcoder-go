package main

import "fmt"

var (
	n, k int
	m    = map[pair]int{}
)

type pair struct{ n, k int }

func f(n, k int) int {
	if _, ok := m[pair{n, k}]; ok {
		return m[pair{n, k}]
	}
	r := n / 10
	if k > 0 {
		r += 1
	}
	for i := 1; i < min(9, n)+1; i++ {
		r += f((n-i)/10, k/i)
	}
	m[pair{n, k}] = r
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Scan(&n, &k)
	fmt.Println(f(n, k) - 1)
}
