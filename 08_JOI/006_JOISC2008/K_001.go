package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 100005

var n int
var tr [N]int

func add(p, q int) {
	for i := p; i <= n; i += i & (-i) {
		tr[i] += q
	}
}

func sum(p int) int {
	res := 0
	for i := p; i >= 1; i -= i & (-i) {
		res += tr[i]
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	type st struct {
		pos, ty, id int
	}

	var m, k int
	fmt.Fscan(in, &n, &m, &k)
	vt := make([]st, 0)
	for i := 1; i <= n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		vt = append(vt, st{a, 1, i})
		vt = append(vt, st{b + 1, -1, i})
	}
	var q, r [N]int
	for i := 1; i <= m; i++ {
		var p int
		fmt.Fscan(in, &p, &q[i], &r[i])
		vt = append(vt, st{p, 3, i})
	}
	sort.Slice(vt, func(p, q int) bool {
		if vt[p].pos != vt[q].pos {
			return vt[p].pos < vt[q].pos
		}
		return vt[p].ty < vt[q].ty
	})
	var ans [N]int
	for _, x := range vt {
		if x.ty < 3 {
			add(x.id, x.ty)
		} else {
			ans[x.id] = sum(r[x.id]) - sum(q[x.id]-1)
		}
	}
	for i := 1; i <= m; i++ {
		fmt.Fprintln(out, ans[i])
	}
}
