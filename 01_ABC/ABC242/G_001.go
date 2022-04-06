package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 1000006

var (
	t   = make([]int, N)
	col = make([]int, N)
)

func update(x, flag int, k *int) {
	if flag != 0 {
		t[col[x]]++
		if t[col[x]]&1 == 0 {
			(*k)++
		}
	} else {
		t[col[x]]--
		if t[col[x]]&1 != 0 {
			(*k)--
		}
	}
	return
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	len := 100
	bl := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &col[i])
		bl[i] = (i-1)/len + 1
	}

	var m int
	fmt.Fscan(in, &m)
	type query struct{ l, r, id int }
	q := make([]query, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &q[i].l, &q[i].r)
		q[i].id = i
	}
	sort.Slice(q, func(i, j int) bool {
		if bl[q[i].l] == bl[q[j].l] {
			if bl[q[i].l]%2 != 0 {
				return q[i].r < q[j].r
			} else {
				return q[i].r > q[j].r
			}
		}
		return bl[q[i].l] < bl[q[j].l]
	})

	l := 1
	r := 1
	tmp := 0
	t[col[1]]++
	ans := make([]int, N)
	for i := 1; i <= m; i++ {
		if q[i].l == q[i].r {
			ans[q[i].id] = 0
			continue
		}
		for r < q[i].r {
			r++
			update(r, 1, &tmp)
		}
		for r > q[i].r {
			update(r, 0, &tmp)
			r--
		}
		for l < q[i].l {
			update(l, 0, &tmp)
			l++
		}
		for l > q[i].l {
			l--
			update(l, 1, &tmp)
		}
		ans[q[i].id] = tmp
	}
	for i := 1; i <= m; i++ {
		fmt.Fprintln(out, ans[i])
	}
}
