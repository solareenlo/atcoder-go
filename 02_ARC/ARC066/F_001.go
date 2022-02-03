package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 300005

var (
	n int
	t int
	w = make([]int, N)
	a = make([]int, N)
	s = make([]int, N)
	g = make([]int, N)
	l = make([]int, N)
	r = make([]int, N)
	f = make([]int, N)
	q = make([]int, N)
)

func sum(n int) int {
	return n * (n + 1) / 2
}

func cal(x, y int) float64 {
	return float64(g[x]-g[y]) / float64(x-y)
}

func pop(x int) {
	for t > 1 && cal(q[t-1], q[t]) <= float64(x) {
		t--
	}
}

func push(x int) {
	for t > 1 && cal(q[t-1], q[t]) <= cal(q[t], x) {
		t--
	}
	t++
	q[t] = x
}

func init_s() {
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + a[i]
	}
}

func INIT(ff []int) {
	t = 1
	q[t] = 0
	init_s()
	for i := 1; i <= n; i++ {
		pop(i)
		ff[i] = max(ff[i-1], ff[q[t]]+sum(i-q[t])-s[i]+s[q[t]])
		g[i] = ff[i] + sum(i-1) + s[i]
		push(i)
	}
}

func work(x, y int) {
	if x == y {
		return
	}
	var i int
	m := (x + y) >> 1
	t = 0
	for i = x; i <= m; i++ {
		g[i] = l[i] + sum(i-1) + s[i]
		push(i)
	}
	for ; i <= y; i++ {
		pop(i)
		w[i] = l[q[t]] + r[i+1] + sum(i-q[t]) - s[i] + s[q[t]]
	}
	for f[y], i = max(f[y], w[y]), y-1; i > m; i-- {
		w[i] = max(w[i], w[i+1])
		f[i] = max(f[i], w[i])
	}
	work(x, m)
	work(m+1, y)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	a = make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		f[i] = -int(1e9)
	}
	INIT(l)
	a = rev(a)
	INIT(r)
	a = rev(a)
	r = rev(r)
	init_s()
	work(0, n)
	a = rev(a)
	f = rev(f)
	l = rev(l)
	r = rev(r)
	l, r = r, l
	init_s()
	work(0, n)
	a = rev(a)
	f = rev(f)
	l = rev(l)
	r = rev(r)
	l, r = r, l

	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var x, k int
		fmt.Fscan(in, &x, &k)
		fmt.Fprintln(out, max(l[x-1]+r[x+1], f[x]+a[x]-k))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rev(a []int) []int {
	tmp := reverseOrderInt(a[1 : n+1])
	for i := 0; i < n; i++ {
		a[i+1] = tmp[i]
	}
	return a
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
