package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, m int
	a, b string
)

func v(t int) int {
	c := 0
	s := [2]int{}
	s1 := [2]int{}
	s2 := [2]int{}
	for i := 0; i < n+m; i++ {
		s1[a[i]-'0']++
		s2[b[i]-'0']++
		if s[t]+1 > max(s1[t], s2[t]) {
			t ^= 1
		} else {
			c++
		}
		s[t]++
	}
	return c
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m, &a, &b)

	fmt.Println(max(v(0), v(1)) - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
