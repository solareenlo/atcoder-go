package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 200005
const mod = 1_000_000_007

func lowbit(x int) int {
	return x & -x
}

type node struct{ l, r int }

var (
	cnt int
	C   int
	n   int
	m   int
	i   int
	fir int
	j   int
	g   = [N]int{}
	nv  int
	c   = [N]int{}
	A   = [N]int{}
	B   = [N]int{}
	f   = [N]node{}
	lsh = [N]node{}
)

func F(x *int) {
	if *x > mod {
		*x -= mod
	}
}

func sum(x int) int {
	ret := 0
	for x > 0 {
		ret += c[x]
		F(&ret)
		x -= lowbit(x)
	}
	return ret
}

func add(x, y int) {
	for x < N {
		c[x] += y
		F(&c[x])
		x += lowbit(x)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i = 1; i <= n; i++ {
		fmt.Fscan(in, &A[i])
	}

	fir = n + 1
	for i = 1; i <= m; i++ {
		fmt.Fscan(in, &B[i])
	}

	for j, i = 1, 1; i <= n; i++ {
		if A[i] > B[1] {
			fir = i
			break
		}
	}

	for i = fir; i <= n; i++ {
		for j <= m && B[j] < A[i] {
			j++
		}
		if j > m {
			break
		}
		cnt++
		f[cnt].l = A[i] - B[j-1]
		f[cnt].r = B[j] - A[i]
	}

	tmp := f[1 : cnt+1]
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].l == tmp[j].l {
			return tmp[i].r > tmp[j].r
		}
		return tmp[i].l < tmp[j].l
	})

	for i = 1; i <= cnt; i++ {
		if f[i].l != f[i-1].l || f[i].r != f[i-1].r {
			C++
			g[C] = f[i].r
		}
	}

	for i = 1; i <= C; i++ {
		lsh[i].r = i
		lsh[i].l = g[i]
	}

	tmp = lsh[1 : C+1]
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].l == tmp[j].l {
			return tmp[i].r > tmp[j].r
		}
		return tmp[i].l < tmp[j].l
	})

	for nv, i = 1, 1; i <= C; i++ {
		tmp := 0
		if lsh[i].l != lsh[i-1].l {
			tmp = 1
		}
		nv += tmp
		g[lsh[i].r] = nv
	}

	add(1, 1)
	for i = 1; i <= C; i++ {
		add(g[i], sum(g[i]-1))
	}

	fmt.Println(sum(nv) % mod)
}
