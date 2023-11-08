package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a [1000000]int

	var n int
	fmt.Fscan(in, &n)
	s := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		s += a[i]
	}
	s /= n
	s1, s2 := 0, 0
	for i := 1; i <= n; i++ {
		if a[i] >= s+1 {
			s1 += a[i] - (s + 1)
		} else {
			s2 += s - a[i]
		}
	}
	fmt.Println(max(s1, s2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
