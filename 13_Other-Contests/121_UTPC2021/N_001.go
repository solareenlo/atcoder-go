package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200005

	var n, m int
	fmt.Fscan(in, &n, &m)
	var p [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}
	o := 1
	var a, b [N]int
	for i := 1; i <= n; i, o = i+1, p[o] {
		a[i] = o
		if b[o] != 0 {
			fmt.Println("No")
			return
		}
		b[o] = i
	}
	A := make([][]pair, N)
	B := make([][]pair, N)
	var e [N]pair
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[i] = pair{u, v}
		u = b[u]
		v = b[v]
		if u > v {
			u, v = v, u
		}
		B[u] = append(B[u], pair{v, i})
		A[v] = append(A[v], pair{u, i})
	}
	st := make([]int, 0)
	for i := 1; i <= n; i++ {
		sortPair(A[i])
		sortPair(B[i])
		for j := len(A[i]) - 1; j >= 0; j-- {
			if st[len(st)-1] != A[i][j].x {
				fmt.Println("No")
				return
			}
			st = st[:len(st)-1]
		}
		for j := 0; j < len(B[i]); j++ {
			st = append(st, i)
		}
	}
	var E [N][]int
	var deg [N]int
	for i := 1; i <= n; i++ {
		o = 0
		for j := 0; j < len(B[i]); j++ {
			u := B[i][j].y
			if o != 0 {
				E[o] = append(E[o], u)
				deg[u]++
			}
			o = u
		}
		for j := 0; j < len(A[i]); j++ {
			u := A[i][j].y
			if o != 0 {
				E[o] = append(E[o], u)
				deg[u]++
			}
			o = u
		}
	}
	q := make([]int, 0)
	for i := 1; i <= m; i++ {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	ans := make([]int, 0)
	for len(q) != 0 {
		u := q[0]
		q = q[1:]
		ans = append(ans, u)
		for _, v := range E[u] {
			deg[v]--
			if deg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	fmt.Println("Yes")
	var vis [N]bool
	for _, i := range ans {
		u := e[i].x
		v := e[i].y
		fmt.Println(u, v)
		p[u], p[v] = p[v], p[u]
	}
	for i := 1; i <= n; i++ {
		if !vis[i] {
			for j := p[i]; j != i; j = p[j] {
				fmt.Println(i, j)
				vis[j] = true
			}
		}
	}
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}
