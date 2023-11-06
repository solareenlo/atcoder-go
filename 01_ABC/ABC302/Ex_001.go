package main

import (
	"bufio"
	"fmt"
	"os"
)

const MXN = 200200

var prt, a, b, Siz, Ans [MXN]int
var Res int
var e [MXN][]int

func find(x int) int {
	if prt[x] == x {
		return x
	}
	return find(prt[x])
}

func dfs(u, fa int) {
	x := find(a[u])
	y := find(b[u])
	lst := Res
	lx := Siz[x]
	ly := Siz[y]
	if x == y {
		if Siz[x] == 0 {
			Siz[x] = 1
			Res++
		}
	} else {
		if Siz[x] == 0 || Siz[y] == 0 {
			Res++
		}
		if Siz[x] > Siz[y] {
			prt[y] = x
			Siz[x] |= Siz[y]
		} else {
			prt[x] = y
			Siz[y] |= Siz[x]
		}
	}
	Ans[u] = Res
	for _, v := range e[u] {
		if v != fa {
			dfs(v, u)
		}
	}
	prt[x] = x
	prt[y] = y
	Siz[x] = lx
	Siz[y] = ly
	Res = lst
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		e[x] = append(e[x], y)
		e[y] = append(e[y], x)
	}
	for i := 1; i < n+1; i++ {
		prt[i] = i
	}
	dfs(1, 0)
	for i := 2; i <= n; i++ {
		fmt.Printf("%d ", Ans[i])
	}
}
