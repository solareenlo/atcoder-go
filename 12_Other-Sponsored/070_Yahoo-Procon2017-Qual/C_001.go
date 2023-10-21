package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, k)
	st := make(map[int]bool)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
		st[a[i]] = true
	}
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}
	var L [100009]int
	for i := 0; i < n; i++ {
		idx := 0
		for j := 0; j < min(len(s[i]), len(s[a[0]])); j++ {
			if s[i][j] == s[a[0]][j] {
				idx++
			} else {
				break
			}
		}
		L[i] = idx
	}
	A, B := int(1e9), -1
	for i := 0; i < n; i++ {
		if st[i] {
			A = min(A, L[i])
		} else {
			B = max(B, L[i])
		}
	}
	if A <= B {
		fmt.Println(-1)
	} else {
		fmt.Println(s[a[0]][:B+1])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
