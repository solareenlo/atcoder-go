package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	s := [2]int{}
	s1 := make([]int, 500000)
	a := make([]int, n+2)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i])
		s[i&1] += a[i]
		s1[i] = s[1] - s[0]
	}

	if n&1 != 0 {
		l := 0
		r := s[1]
		for l < r {
			m := (l + r + 1) / 2
			mn := 0
			for i := 1; i < (n/2)+1; i++ {
				if s1[i*2-1]-mn >= m {
					mn = min(mn, s1[i*2])
				}
			}
			if s1[n]-mn >= m {
				l = m
			} else {
				r = m - 1
			}
		}
		fmt.Println(s[0]+l, s[1]-l)
	} else {
		fmt.Println(max(s[0], s[1]), min(s[0], s[1]))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
