package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 500500

var (
	n int
	m int
	L = make([]int, N)
	a = make([]int, N)
	c = make([]int, N)
)

func Dich(l, r int) {
	if l == r {
		return
	}
	mid := (l + r) >> 1
	ql := l
	p := l
	Dich(l, mid)
	Dich(mid+1, r)
	L[mid+1] = n + 1
	for i := mid; i >= l; i-- {
		L[i] = min(L[i+1], a[i]-m)
	}
	for i := mid + 1; i <= r; i++ {
		for L[p] < a[i] {
			c[ql] = a[p]
			ql++
			p++
		}
		c[ql] = a[i]
		ql++
	}
	for p <= mid {
		c[ql] = a[p]
		ql++
		p++
	}
	for i := l; i <= r; i++ {
		a[i] = c[i]
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		a[x] = i
	}

	Dich(1, n)

	b := make([]int, N)
	for i := 1; i <= n; i++ {
		b[a[i]] = i
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, b[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
