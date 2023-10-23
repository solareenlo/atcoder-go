package main

import "fmt"

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)
	s1, c := 0, 0
	n := len(S)
	for i := 0; i < n; i++ {
		if S[i] == 'A' {
			c++
			s1 += c
		} else {
			c = 0
		}
	}
	if c == n {
		fmt.Println(c * N * (c*N + 1) / 2)
		return
	}
	s2 := 0
	for i := 0; i < n; i++ {
		if S[i] == 'A' {
			c++
			s2 += c
		} else {
			c = 0
		}
	}
	fmt.Println(s1 + s2*(N-1))
}
