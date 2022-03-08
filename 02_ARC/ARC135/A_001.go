package main

import "fmt"

var m = map[int]int{1: 1, 2: 2, 3: 3, 4: 4}

func f(n int) int {
	if _, ok := m[n]; ok {
		return m[n]
	}
	m[n] = f(n/2) * f((n+1)>>1) % 998244353
	return m[n]
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(f(x))
}
