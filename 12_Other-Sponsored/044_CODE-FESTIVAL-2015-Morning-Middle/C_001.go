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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	l := 0
	r := n - 1
	cnt := 0
	for t := 0; t < n>>1; t++ {
		cl := a[l]*2 + a[l+1]
		cr := a[r]*2 + a[r-1]
		cnt += min(cl, cr) + 1
		if cl <= cr {
			l += 2
			a[l] += cl - a[l-2] + 2
		} else {
			r -= 2
			a[r] += cr - a[r+2] + 2
		}
	}
	fmt.Println(cnt)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
