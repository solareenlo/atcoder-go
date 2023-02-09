package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, ans int
var fa, g, d []int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 200005

	var q int
	fmt.Fscan(in, &n, &q)

	fa = make([]int, N)
	g = make([]int, N)
	d = make([]int, N)
	ans = n * n
	for i := 0; i < n; i++ {
		fa[i] = i
		g[i] = n
	}
	for i := 1; i <= q; i++ {
		var a, b, c, e int
		fmt.Fscan(in, &a, &b, &c, &e)
		add((a-b+n)%n, (c-e+n)%n, (e-b+n)%n)
		fmt.Fprintln(out, ans)
	}
}

func add(x, y, z int) {
	find(x)
	find(y)
	z = (d[x] + z - d[y] + n) % n
	x = fa[x]
	y = fa[y]
	if x == y {
		ans -= g[x]
		g[x] = gcd(g[x], z)
		ans += g[x]
	} else {
		fa[y] = x
		d[y] = z
		ans -= g[x] + g[y]
		g[x] = gcd(g[x], g[y])
		ans += g[x]
	}
}

func find(k int) int {
	if k == fa[k] {
		return k
	}
	f := find(fa[k])
	d[k] = (d[k] + d[fa[k]]) % n
	fa[k] = f
	return fa[k]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
