package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	const N = 100005
	x := make([]int, N)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &x[i])
	}

	var m, kk int
	fmt.Fscan(in, &m, &kk)
	a := make([]int, N)
	b := make([]int, N)
	for i := 1; i < n; i++ {
		a[i] = i
		b[i] = i
	}

	for i := 1; i < m+1; i++ {
		var y int
		fmt.Fscan(in, &y)
		a[y-1], a[y] = a[y], a[y-1]
	}

	c := make([]int, N)
	for ; kk > 0; kk >>= 1 {
		if kk&1 != 0 {
			for i := 1; i < n; i++ {
				b[i] = a[b[i]]
			}
		}
		for i := 1; i < n; i++ {
			c[i] = a[a[i]]
		}
		for i := 1; i < n; i++ {
			a[i] = c[i]
		}
	}

	ans := x[1]
	for i := 1; i < n+1; i++ {
		fmt.Fprint(out, ans, ".0\n")
		ans += x[b[i]+1] - x[b[i]]
	}
}
