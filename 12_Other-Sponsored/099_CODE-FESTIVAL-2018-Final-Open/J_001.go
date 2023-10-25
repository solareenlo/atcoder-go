package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type wolf struct {
	a, b, c int
	d       []string
}

var ev []wolf

func awoo(t []int) {
	n := len(t)
	for i := n / 2; i >= 1; i /= 2 {
		z := make([]int, n)
		g := make([]int, 0)
		for j := 0; j < n; j++ {
			if (t[j] & i) != 0 {
				z[j] = 1
			}
		}
		for j := 0; j < n; j++ {
			if z[j] == 0 {
				g = append(g, j)
			}
		}
		tar := make([]int, 0)
		for j := 0; j < n; j++ {
			if (j & i) == 0 {
				tar = append(tar, j)
			}
		}
		for j := 1; j <= i; j *= 2 {
			var L1, L2 wolf
			L1.a = len(ev)
			L1.b = len(ev)
			L2.b = len(ev)
			L2.a = len(ev) + 1
			L1.c = j
			L2.c = -j
			L1.d = strings.Split(strings.Repeat("x", n), "")
			L2.d = strings.Split(strings.Repeat("x", n), "")
			for k := 0; k < n/2; k++ {
				if (j & (g[k] ^ tar[k])) != 0 {
					L1.d[((g[k])&(^j))+j] = "y"
					L2.d[(g[k])&(^j)] = "y"
					g[k] ^= j
				}
			}
			ev = append(ev, L1)
			ev = append(ev, L2)
			to := make([]int, n)
			for k := 0; k < n; k++ {
				to[k] = t[k]
				if L1.d[k] == "y" {
					to[k] = t[k-j]
				}
				if L2.d[k] == "y" {
					to[k] = t[k+j]
				}
			}
			t = to
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a int
	fmt.Fscan(in, &a)
	x := make([]int, a)
	y := make([]int, a)
	for i := 0; i < a; i++ {
		x[i] = -1
		y[i] = -1
	}
	p := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Fscan(in, &p[i])
		p[i]--
		x[p[i]] = i
	}
	w := make([]pair, a)
	q := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Fscan(in, &q[i])
		q[i]--
		y[q[i]] = i
		w[i] = pair{q[i], i}
	}
	sortPair(w)
	for i := 0; i < a; i++ {
		if ^y[i] != 0 && ^x[i] == 0 {
			fmt.Fprintln(out, -1)
			return
		}
	}
	r := make([]int, a)
	for i := 0; i < a; i++ {
		r[i] = q[i]
	}
	sort.Ints(r)
	st := make([]int, a)
	for i := range st {
		st[i] = -1
	}
	con := make([]int, a)
	for i := a - 1; i >= 0; i-- {
		con[i] = 1
		if i < a-1 && r[i] == r[i+1] {
			con[i] = con[i+1] + 1
		}
	}
	used := make([]bool, a)
	for i := 0; i < a; i++ {
		if i == 0 || r[i] != r[i-1] {
			st[x[r[i]]] = i
			used[x[r[i]]] = true
		}
	}
	at := 0
	for i := 0; i < a; i++ {
		if i == 0 || r[i] != r[i-1] {
			continue
		}
		for used[at] {
			at++
		}
		st[at] = i
		at++
	}
	awoo(st)
	for i := 1; i < a; i *= 2 {
		id := len(ev)
		var ad wolf
		ad.a = id
		ad.b = id
		ad.c = i
		ad.d = strings.Split(strings.Repeat("x", a), "")
		for j := 0; j < a; j++ {
			if j-i >= 0 && r[j] == r[j-i] {
				ad.d[j] = "y"
			}
		}
		ev = append(ev, ad)
	}

	for i := 0; i < a; i++ {
		st[i] = w[i].y
	}
	awoo(st)
	fmt.Fprintln(out, len(ev))
	for i := 0; i < len(ev); i++ {
		fmt.Fprintln(out, ev[i].a, ev[i].b, ev[i].c, strings.Join(ev[i].d, ""))
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
