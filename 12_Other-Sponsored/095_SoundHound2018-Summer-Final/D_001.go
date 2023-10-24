package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100003

var fa, siz, pas, r [N]int

func find(x int) int {
	for x != fa[x] {
		x = fa[x]
	}
	return x
}

func unionset(u, v int) {
	x := find(u)
	y := find(v)
	if x == y {
		return
	}
	if siz[x] < siz[y] {
		x, y = y, x
	}
	fa[y] = x
	siz[x] += siz[y]
	pas[y] = r[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b [30]int
	un := make([]map[int]int, N)
	for i := range un {
		un[i] = make(map[int]int)
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fa[i] = i
		siz[i] = 1
	}
	for m > 0 {
		m--
		var opt, x, y int
		fmt.Fscan(in, &opt, &x, &y)
		if opt == 1 {
			unionset(x, y)
			if x > y {
				x, y = y, x
			}
			un[x][y]++
		} else if opt == 2 {
			r[find(x)]++
		} else {
			t1, t2, ans := 0, 0, 0
			if x > y {
				x, y = y, x
			}
			if find(x) != find(y) {
				fmt.Println("No")
				continue
			}
			if _, ok := un[x][y]; ok {
				fmt.Println("Yes")
				continue
			}
			var i int
			for i = x; i != fa[i]; i = fa[i] {
				t1++
				a[t1] = i
			}
			t1++
			a[t1] = i
			for i := y; i != fa[i]; i = fa[i] {
				t2++
				b[t2] = i
			}
			t2++
			b[t2] = i
			var j int
			for i, j = t1, t2; a[i] == b[j]; i, j = i-1, j-1 {
				ans += r[a[i]] - pas[a[i]]
			}
			ans -= max(pas[a[i]], pas[b[j]])
			if ans != 0 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
