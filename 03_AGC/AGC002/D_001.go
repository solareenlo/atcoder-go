package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100100

type Node struct {
	x, y, z, l, r, las, ord int
}

var (
	q  = make([]Node, N)
	fa = make([]int, N)
)

func find_fa(a int) int {
	if fa[a] == a {
		return a
	}
	fa[a] = find_fa(fa[a])
	return fa[a]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}

	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var x, y, z int
		fmt.Fscan(in, &x, &y, &z)
		q[i] = Node{x, y, z, 0, m - 1, m, i}
	}

	siz := make([]int, n+1)
	for l := 0; m>>l > 0; l++ {
		for i := 1; i <= n; i++ {
			fa[i] = i
			siz[i] = 1
		}
		sl := -1
		sr := -1
		beg := 0
		i := 0
		var fal, far int
		for j := 0; j < Q; j++ {
			if q[j].l != sl || q[j].r != sr {
				beg = j
				sl = q[j].l
				sr = q[j].r
			}
			mid := (q[j].l + q[j].r + 1) >> 1
			for i < mid && i+1 < m {
				fal = find_fa(a[i])
				far = find_fa(b[i])
				if fal != far {
					siz[far] += siz[fa[fal]]
					fa[fal] = far
				}
				i++
			}
			fal = find_fa(q[j].x)
			far = find_fa(q[j].y)
			tmp := 0
			if fal == far {
				tmp = 1
			}
			if siz[fal]+siz[far]-tmp*siz[fal] >= q[j].z {
				q[j].r = mid - 1
				q[j].las = mid
				q[j], q[beg] = q[beg], q[j]
				beg++
			} else {
				q[j].l = mid + 1
			}
		}
	}

	ans := make([]int, N)
	for i := 0; i < Q; i++ {
		ans[q[i].ord] = q[i].las
	}
	for i := 0; i < Q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}
