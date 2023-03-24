package main

import (
	"bufio"
	"fmt"
	"os"
)

var w, W, h int
var x, b, idx [110]int
var y [110][70]int
var prep []M

func main() {
	in := bufio.NewReader(os.Stdin)

	var r, a int
	fmt.Fscan(in, &w, &h, &r, &a)
	r--

	for i := 0; i < a; i++ {
		fmt.Fscan(in, &x[i], &b[i])
		x[i]--
		for j := 0; j < b[i]; j++ {
			fmt.Fscan(in, &y[i][j])
			y[i][j]--
		}
		idx[i] = i
		for j := i - 1; j >= 0; j-- {
			if x[idx[j+1]] < x[idx[j]] {
				tmp := idx[j+1]
				idx[j+1] = idx[j]
				idx[j] = tmp
			}
		}
	}

	if h == 1 {
		t := NewM(0)
		for i := 0; i < w; i++ {
			if i != 0 {
				t.flip(i-1, i)
			}
			t.flip(i, i)
			if i != w-1 {
				t.flip(i+1, i)
			}
		}
		u := make([]int, 70)
		for i := 0; i < b[0]; i++ {
			u[y[0][i]] = 1
		}
		solve(t, u, 70)
		Print(u)
		return
	}

	W = 2 * w
	t := NewM(0)
	for i := 0; i < w; i++ {
		if i != 0 {
			t.flip(i-1, i)
		}
		t.flip(i, i)
		if i != w-1 {
			t.flip(i+1, i)
		}
	}

	ini := NewM(0)
	fi := NewM(0)
	for i := range t.a {
		copy(ini.a[i], t.a[i])
		copy(fi.a[i], t.a[i])
	}
	for i := w; i < W; i++ {
		fi.flip(i-w, i)
		t.flip(i-w, i)
	}
	for i := 0; i < w; i++ {
		ini.flip(i+w, i)
		t.flip(i+w, i)
	}
	ini.flip(W, w)
	t.flip(W, W)
	fi.flip(w, W)
	if x[idx[0]] == 0 {
		for j := 0; j < b[idx[0]]; j++ {
			ini.flip(y[idx[0]][j], w)
		}
	}

	prep = make([]M, 70)
	for i := range prep {
		prep[i] = NewM(0)
	}
	for i := range t.a {
		copy(prep[0].a[i], t.a[i])
	}
	for i := 0; i < 61; i++ {
		prep[i+1] = mul(prep[i], prep[i])
	}
	f := NewM(1)
	f2 := NewM(0)
	ph := 1
	for i := 0; i < a; i++ {
		th := x[idx[i]]
		if th == 0 {
			continue
		}
		if th > r && r >= ph {
			f2 = mul(pow(r-ph), f)
		}
		f = mul(pow(th-ph), f)
		tm := NewM(1)
		for j := 0; j < b[idx[i]]; j++ {
			tw := y[idx[i]][j]
			tm.flip(tw+w, W)
		}
		f = mul(tm, f)
		ph = th
	}
	if r >= ph {
		f2 = mul(pow(r-ph), f)
	}
	f = mul(pow(h-1-ph), f)

	z := NewM(0)
	for i := 0; i <= W; i++ {
		for k := 0; k <= W; k++ {
			for j := 0; j <= w; j++ {
				if f.get(i, k) != 0 && ini.get(k, j) != 0 {
					z.flip(i, j)
				}
			}
		}
	}

	A := NewM(0)
	for i := 0; i <= w; i++ {
		for k := 0; k <= W; k++ {
			for j := 0; j <= w; j++ {
				if fi.get(i, k) != 0 && z.get(k, j) != 0 {
					A.flip(i, j)
				}
			}
		}
	}

	u := make([]int, 124)
	for i := 0; i < w; i++ {
		if A.get(i, w) != 0 {
			u[i] ^= 1
		}
	}

	solve(A, u, 124)
	u[w] = 1
	if r == 0 {
		Print(u)
		return
	}

	tu := make([]int, 124)
	for i := 0; i <= W; i++ {
		for j := 0; j <= w; j++ {
			tu[i] ^= ini.get(i, j) & u[j]
		}
	}
	for i := 0; i <= W; i++ {
		u[i] = tu[i]
		tu[i] = 0
	}
	for i := 0; i <= W; i++ {
		for j := 0; j <= W; j++ {
			tu[i] ^= f2.get(i, j) & u[j]
		}
	}
	Print(tu)
}

func solve(m M, u []int, n int) {
	for i := 0; i < w; i++ {
		pv := -1
		for j := i; j < w; j++ {
			if m.get(j, i) != 0 {
				pv = j
				break
			}
		}
		for j := 0; j < w; j++ {
			tmp := m.get(i, j)
			m.set(i, j, m.get(pv, j))
			m.set(pv, j, tmp)
		}
		if pv != -1 {
			u[i], u[pv] = u[pv], u[i]
		}
		for j := 0; j < w; j++ {
			if i == j || m.get(j, i) == 0 {
				continue
			}
			for k := 0; k < w; k++ {
				if m.get(i, k) != 0 {
					m.flip(j, k)
				}
			}
			if j < n && i < n {
				u[j] ^= u[i]
			}
		}
	}
}

type M struct {
	a [][]uint64
}

func NewM(d int) M {
	var m M
	m.a = make([][]uint64, 128)
	for i := range m.a {
		m.a[i] = make([]uint64, 2)
	}
	for i := 0; i < 128; i++ {
		m.set(i, i, d)
	}
	return m
}

func (m *M) set(x, y, v int) {
	p := y / 64
	q := y % 64
	if x < 0 || 127 < x || p < 0 || 1 < p {
		return
	}
	if v != 0 {
		m.a[x][p] |= 1 << q
	} else {
		m.a[x][p] = m.a[x][p] & (^(1 << q))
	}
}

func (m M) get(x, y int) int {
	p := y / 64
	q := y % 64
	if x < 0 || 127 < x || p < 0 || 1 < p {
		return 0
	}
	return int(m.a[x][p]>>q) & 1
}

func (m *M) flip(x, y int) {
	p := y / 64
	q := y % 64
	if x < 0 || 127 < x || p < 0 || 1 < p {
		return
	}
	m.a[x][p] = m.a[x][p] ^ (1 << q)
}

func mul(ma, mb M) M {
	res := NewM(0)
	for j := 0; j <= W; j++ {
		var z [2]int
		for k := 0; k <= W; k++ {
			if mb.get(k, j) != 0 {
				z[k/64] |= 1 << (k % 64)
			}
		}
		for i := 0; i < 127; i++ {
			s := 0
			u := uint64(z[0]) & ma.a[i][0]
			tmp0 := parity(u >> 32)
			tmp1 := parity(u)
			s ^= tmp0
			s ^= tmp1
			u = uint64(z[1]) & ma.a[i][1]
			s ^= parity(u >> 32)
			s ^= parity(u)
			if s != 0 {
				res.flip(i, j)
			}
		}
	}
	return res
}

func pow(p int) M {
	res := NewM(1)
	for i := 0; i < 61; i++ {
		if ((p >> i) & 1) != 0 {
			res = mul(res, prep[i])
		}
	}
	return res
}

func parity(n uint64) int {
	parity := 0
	for n > 0 {
		if parity == 0 {
			parity = 1
		} else {
			parity = 0
		}
		n = n & (n - 1)
	}
	return parity
}

func Print(u []int) {
	c := 0
	for i := 0; i < w; i++ {
		c += u[i]
	}
	fmt.Printf("%d", c)
	for i := 0; i < w; i++ {
		if u[i] != 0 {
			fmt.Printf(" %d", i+1)
		}
	}
	fmt.Printf("\n")
}
