package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 500005

	type node struct {
		l, r, id, pre, nxt int
	}

	var n, T int
	fmt.Fscan(in, &n, &T)

	var a [N]node
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i].l, &a[i].r)
		a[i].id = i
	}
	tmp := a[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].l < tmp[j].l
	})
	a[0].l = T
	a[0].r = T
	a[0].id = 0
	a[0].nxt = 1
	a[1].pre = 1
	r := a[1].r
	fst := 1
	lst := 1
	for i := 2; i <= n; i++ {
		if a[i].r >= r {
			a[i].pre = lst
			a[lst].nxt = i
			r = a[i].r
			lst = i
		} else {
			a[i].pre = 0
			a[i].nxt = fst
			a[fst].pre = i
			a[0].nxt = i
			fst = i
		}
	}

	pos := 0
	ans := 0
	for i := 1; i <= n; i++ {
		nxt := a[pos].nxt
		if a[pos].r > a[nxt].l {
			ans += a[pos].r - a[nxt].l
		}
		pos = nxt
	}
	fmt.Println(ans)
	pos = a[0].nxt
	for i := 1; i <= n; i++ {
		if i == n {
			fmt.Println(a[pos].id)
		} else {
			fmt.Printf("%d ", a[pos].id)
		}
		pos = a[pos].nxt
	}
}
