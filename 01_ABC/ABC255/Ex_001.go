package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Query struct {
	id         int
	Qt, Ql, Qr int64
}
type Number struct {
	oldv, newv int64
	id         int
}

var Qu [200000]Query
var wei [400000]int64
var tree [1600000]int64
var tim [1600000]int64
var sum [1600000]int

const mod = 998244353

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int64
	var q int
	fmt.Fscan(in, &n, &q)
	num := make([]Number, 400000)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &Qu[i].Qt, &Qu[i].Ql, &Qu[i].Qr)
		Qu[i].Ql--
		Qu[i].id = i
		num[i].oldv = Qu[i].Ql
		num[i].id = i
		num[i+q].oldv = Qu[i].Qr
		num[i+q].id = i + q
	}
	sort.Slice(num[:q+q], func(i, j int) bool {
		return num[i].oldv < num[j].oldv
	})

	cnt := 0
	for i := 0; i < q+q; i++ {
		if i == 0 || num[i].oldv > num[i-1].oldv {
			x := int64(1)
			y := num[i].oldv
			if i > 0 {
				x = num[i-1].oldv + 1
			}
			wei[cnt] = (((((x + y) % mod) * ((y - x + 1) % mod)) % mod) * ((mod + 1) / 2)) % mod
			num[i].newv = int64(cnt)
			cnt++
		}
		if num[i].id < q {
			Qu[num[i].id].Ql = int64(cnt)
		} else {
			Qu[num[i].id-q].Qr = int64(cnt) - 1
		}
	}
	build(1, 0, cnt-1)
	for i := 0; i < q; i++ {
		fmt.Fprintln(out, modify(1, 0, int64(cnt-1), Qu[i].Ql, Qu[i].Qr, Qu[i].Qt%mod))
	}
}

func build(i, l, r int) {
	tim[i] = -1
	sum[i] = 0
	if l == r {
		tree[i] = wei[l]
	} else {
		m := (l + r) >> 1
		build(i<<1, l, m)
		build(i<<1|1, m+1, r)
		tree[i] = (tree[i<<1] + tree[i<<1|1]) % mod
	}
	return
}

func modify(i, l, r, ql, qr, qt int64) int {
	if r < ql || qr < l {
		return 0
	}
	if ql <= l && r <= qr {
		tim[i] = qt
		t := sum[i]
		sum[i] = int((tree[i] * qt) % mod)
		return (sum[i] - t + mod) % mod
	}
	if tim[i] != -1 {
		tim[i<<1|1] = tim[i]
		tim[i<<1] = tim[i<<1|1]
		sum[i<<1] = int((tim[i] * tree[i<<1]) % mod)
		sum[i<<1|1] = int((tim[i] * tree[i<<1|1]) % mod)
		tim[i] = -1
	}
	m := (l + r) >> 1
	res := 0
	res = (res + modify(i<<1, l, m, ql, qr, qt)) % mod
	res = (res + modify(i<<1|1, m+1, r, ql, qr, qt)) % mod
	sum[i] = (sum[i<<1] + sum[i<<1|1]) % mod
	return res
}
