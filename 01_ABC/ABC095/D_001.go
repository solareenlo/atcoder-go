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
	x := make([]int, n)
	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &v[i])
	}

	a := make([]int, n+1)
	b := make([]int, n+1)
	c := make([]int, n+1)
	d := make([]int, n+1)
	l, r := 0, 0
	for i := 0; i < n; i++ {
		l += v[i]
		r += v[n-i-1]
		a[i+1] = max(a[i], l-x[i])
		b[i+1] = max(b[i], l-x[i]*2)
		c[n-i-1] = max(c[n-i], r-(k-x[n-i-1]))
		d[n-i-1] = max(d[n-i], r-(k-x[n-i-1])*2)
	}

	res := -(1 << 60)
	for i := 0; i <= n; i++ {
		res = max(res, max(a[i]+d[i], b[i]+c[i]))
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
