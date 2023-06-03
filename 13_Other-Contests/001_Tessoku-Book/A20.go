package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s, t string
	fmt.Fscan(in, &s, &t)
	ns := len(s)
	nt := len(t)
	A := make([][]int, ns+1)
	for i := range A {
		A[i] = make([]int, nt+1)
	}
	for i := 1; i <= ns; i++ {
		for j := 1; j <= nt; j++ {
			tmp := 0
			if s[i-1] == t[j-1] {
				tmp = 1
			}
			A[i][j] = max(A[i-1][j], A[i][j-1], A[i-1][j-1]+tmp)
		}
	}
	fmt.Println(A[ns][nt])
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
