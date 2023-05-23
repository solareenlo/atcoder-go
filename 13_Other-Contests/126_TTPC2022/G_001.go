package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	l := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &r[i])
	}
	cl := make([]int, 0)
	cr := make([]int, 0)
	for i := 0; i < n; i++ {
		var check func(int, int, int) bool
		check = func(i, j, k int) bool {
			d1 := j - i
			d2 := k - j
			return d2*l[i]+d1*l[k] < (d1+d2)*l[j]
		}
		for len(cl) >= 2 && !check(cl[len(cl)-2], cl[len(cl)-1], i) {
			cl = cl[:len(cl)-1]
		}
		cl = append(cl, i)
	}
	for i := 0; i < n; i++ {
		var check func(int, int, int) bool
		check = func(i, j, k int) bool {
			d1 := j - i
			d2 := k - j
			return d2*r[i]+d1*r[k] > (d1+d2)*r[j]
		}
		for len(cr) >= 2 && !check(cr[len(cr)-2], cr[len(cr)-1], i) {
			cr = cr[:len(cr)-1]
		}
		cr = append(cr, i)
	}

	ev := make([]pair, 0)
	sz := len(cl)
	for idx := 0; idx < sz-1; idx++ {
		i := cl[idx]
		j := cl[idx+1]
		dy := l[j] - l[i]
		dx := j - i
		var t int
		if dy <= 0 {
			t = dy / dx
		} else {
			t = (dy + dx - 1) / dx
		}
		ev = append(ev, pair{t, 0})
	}
	sz = len(cr)
	for idx := 0; idx < sz-1; idx++ {
		i := cr[idx]
		j := cr[idx+1]
		dy := r[j] - r[i]
		dx := j - i
		var t int
		if dy <= 0 {
			t = dy / dx
		} else {
			t = (dy + dx - 1) / dx
		}
		ev = append(ev, pair{t, 1})
	}
	ev = append(ev, pair{1_000_000_000_001, 0})
	sortPair(ev)
	from := -1_000_000_000_001
	idxl := len(cl) - 1
	idxr := 0
	ans := 0
	inv2 := int(mint(2).inv())
	for _, tmp := range ev {
		to := tmp.x
		lr := tmp.y
		il := cl[idxl]
		ir := cr[idxr]
		tl := from
		tr := to
		if il < ir {
			dy := r[ir] - l[il]
			dx := ir - il
			var t int
			if dy >= 0 {
				t = dy / dx
			} else {
				t = (dy - (dx - 1)) / dx
			}
			tr = min(tr, t+1)
		} else if ir < il {
			dy := l[il] - r[ir]
			dx := il - ir
			var t int
			if dy <= 0 {
				t = dy / dx
			} else {
				t = (dy + dx - 1) / dx
			}
			tl = max(tl, t)
		}
		if tl < tr {
			ans = (ans + (((tr%MOD)-(tl%MOD)+MOD)%MOD)*(((r[ir]%MOD)-(l[il]%MOD)+1+MOD)%MOD)%MOD) % MOD
			ans = (ans + (((((il%MOD)-(ir%MOD)+MOD)%MOD)*(((tr%MOD)+(tl%MOD)-1+MOD)%MOD)%MOD)*(((tr%MOD)-(tl%MOD)+MOD)%MOD)%MOD)*inv2%MOD) % MOD
		}
		from = to
		if lr == 0 {
			idxl--
		} else {
			idxr++
		}
	}
	fmt.Println(ans)
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

type mint int

func (m mint) pow(p int) mint {
	return powMod(m, p)
}

func (m mint) inv() mint {
	return invMod(m)
}

func (m mint) div(n mint) mint {
	return divMod(m, n)
}

const MOD = 998244353

func powMod(a mint, n int) mint {
	res := mint(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func invMod(a mint) mint {
	return powMod(a, MOD-2)
}

func divMod(a, b mint) mint {
	ret := a * modInv(b)
	ret %= MOD
	return ret
}

func modInv(a mint) mint {
	b, u, v := mint(MOD), mint(1), mint(0)
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= MOD
	if u < 0 {
		u += MOD
	}
	return u
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
