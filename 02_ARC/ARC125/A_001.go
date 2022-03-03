package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	s := make([]int, n)
	t := make([]int, m)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}
	for i := range t {
		fmt.Fscan(in, &t[i])
	}

	mdf := n
	for i := 0; i < n; i++ {
		if s[i] != s[0] {
			mdf = min(mdf, min(i, n-i))
		}
	}

	chs := 0
	if t[0] != s[0] {
		chs++
	}
	for i := 0; i < m-1; i++ {
		if t[i] != t[i+1] {
			chs++
		}
	}

	if chs == 0 {
		fmt.Println(m)
	} else if mdf >= n {
		fmt.Println(-1)
	} else {
		fmt.Println(mdf + chs + m - 1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
