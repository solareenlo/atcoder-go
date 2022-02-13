package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 100005

type edge struct{ v, nxt int }

var (
	ed int
	B  int
	e  = make([]edge, N<<1)
	g  = make([]int, N)
	f  = make([]int, N)
	q  = make([]int, N)
)

func adde(x, y int) {
	ed++
	e[ed].v = y
	e[ed].nxt = g[x]
	g[x] = ed
}

func dfs(x, y int) bool {
	var i int
	f[x] = 0
	for i = g[x]; i > 0; i = e[i].nxt {
		if e[i].v != y && !dfs(e[i].v, x) {
			return false
		}
	}
	t := 0
	for i = g[x]; i > 0; i = e[i].nxt {
		if e[i].v != y {
			t++
			q[t] = f[e[i].v] + 1
		}
	}
	if t == 0 {
		return true
	}
	if ^t&1 != 0 {
		t++
		q[t] = 0
	}
	tmp := q[1 : t+1]
	sort.Ints(tmp)
	if q[t] > B {
		return false
	}
	l := 1
	r := t + 1
	for l < r {
		mid := (l + r) >> 1
		l1 := 0
		r1 := t + 1
		for i = 1; i <= t/2; i++ {
			l1++
			if l1 == mid {
				l1++
			}
			r1--
			if r1 == mid {
				r1--
			}
			if q[l1]+q[r1] > B {
				break
			}
		}
		if i <= t/2 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l > t {
		return false
	}
	f[x] = q[l]
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	deg := make([]int, N)
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		adde(x, y)
		adde(y, x)
		deg[x]++
		deg[y]++
	}
	tot := 0
	root := 0
	for i := 1; i <= n; i++ {
		if deg[i]&1 != 0 {
			tot++
		}
		if deg[i] == 1 {
			root = i
		}
	}
	tot >>= 1
	l := 1
	r := n
	for l < r {
		mid := (l + r) >> 1
		B = mid
		if !dfs(root, root) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	fmt.Println(tot, l)
}
