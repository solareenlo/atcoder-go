package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const NN = 200000
const MM = 200000
const INF = 1061109567

var Rand_ = rand.New(rand.NewSource(time.Now().UnixNano()))

var II, ii_, JJ, jj_, ww [MM]int
var bridge [MM]bool
var ds, ww_ []int

func Sort(hh []int, l, r int) {
	for l < r {
		i := l
		j := l
		k := r
		h := hh[l+Rand_.Int()%(r-l)]
		for j < k {
			if ww[hh[j]] == ww[h] {
				j++
			} else if ww[hh[j]] < ww[h] {
				hh[i], hh[j] = hh[j], hh[i]
				i++
				j++
			} else {
				k--
				hh[j], hh[k] = hh[k], hh[j]
			}
		}
		Sort(hh, l, i)
		l = k
	}
}

func find(i int) int {
	if ds[i] < 0 {
		return i
	}
	return find(ds[i])
}

func join(i, j, w int) bool {
	i = find(i)
	j = find(j)
	if i == j {
		return false
	}
	if ds[i] > ds[j] {
		ds[i] = j
		ww_[i] = w
	} else {
		if ds[i] == ds[j] {
			ds[i]--
		}
		ds[j] = i
		ww_[j] = w
	}
	return true
}

func query(i, j int) int {
	w := -1
	for i != j {
		if ww_[i] < ww_[j] {
			w = ww_[i]
			i = ds[i]
		} else {
			w = ww_[j]
			j = ds[j]
		}
	}
	return w
}

var eh [NN][]int
var eo []int
var eo_ [NN]int

func Append(i, h int) {
	o := eo[i]
	eo[i]++
	if o == eo_[i] {
		eo_[i] *= 2
		resize(&eh[i], eo_[i])
	}
	eh[i][o] = h
}

func resize(a *[]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
}

var ta, tb [NN]int
var t int

func dfs1(f, i int) {
	tb[i] = t
	ta[i] = t
	t++
	for o := eo[i]; o > 0; {
		o--
		h := eh[i][o]
		j := i ^ ii_[h] ^ jj_[h]
		if h != f {
			if ta[j] == 0 {
				dfs1(h, j)
				if ta[j] == tb[j] {
					bridge[h] = true
				}
				tb[i] = min(tb[i], tb[j])
			} else {
				tb[i] = min(tb[i], ta[j])
			}
		}
	}
}

func dfs2(p, i int) {
	ta[i] = t
	t++
	for o := eo[i]; o > 0; {
		o--
		h := eh[i][o]
		j := i ^ II[h] ^ JJ[h]
		if j != p {
			dfs2(i, j)
		}
	}
	tb[i] = t
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	hh := make([]int, MM)
	var tree [MM]bool

	var n, m int
	fmt.Fscan(in, &n, &m)
	for h := 0; h < m; h++ {
		fmt.Fscan(in, &II[h], &JJ[h], &ww[h])
		II[h]--
		JJ[h]--
		hh[h] = h
	}
	Sort(hh, 0, m)
	for i := 0; i < n; i++ {
		eo_[i] = 2
		eh[i] = make([]int, eo_[i])
	}
	ds = make([]int, n)
	for i := range ds {
		ds[i] = -1
	}
	ww_ = make([]int, n)
	for i := range ww_ {
		ww_[i] = INF
	}
	var h_ int
	eo = make([]int, NN)
	for h := 0; h < m; h = h_ {
		h_ = h + 1
		for h_ < m && ww[hh[h_]] == ww[hh[h]] {
			h_++
		}
		for g := h; g < h_; g++ {
			h1 := hh[g]
			i := find(II[h1])
			ii_[h1] = i
			j := find(JJ[h1])
			jj_[h1] = j
			eo[i] = 0
			eo[j] = 0
			ta[i] = 0
			ta[j] = 0
		}
		for g := h; g < h_; g++ {
			h1 := hh[g]
			i := ii_[h1]
			j := jj_[h1]
			if i != j {
				Append(i, h1)
				Append(j, h1)
			}
		}
		t = 1
		for g := h; g < h_; g++ {
			h1 := hh[g]
			i := ii_[h1]
			if ta[i] == 0 {
				dfs1(-1, i)
			}
		}
		for g := h; g < h_; g++ {
			h1 := hh[g]
			tree[h1] = join(II[h1], JJ[h1], ww[h1])
		}
	}
	for i := 0; i < n; i++ {
		eo[i] = 0
	}
	for h := 0; h < m; h++ {
		if tree[h] {
			Append(II[h], h)
			Append(JJ[h], h)
		}
	}
	t = 1
	dfs2(-1, 0)
	for h := 0; h < m; h++ {
		if tree[h] && ta[II[h]] > ta[JJ[h]] {
			II[h], JJ[h] = JJ[h], II[h]
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var h, i, j int
		fmt.Fscan(in, &h, &i, &j)
		h--
		i--
		j--
		if !bridge[h] {
			fmt.Fprintln(out, 0)
			continue
		}
		k := JJ[h]
		if (ta[k] <= ta[i] && ta[i] < tb[k]) == (ta[k] <= ta[j] && ta[j] < tb[k]) {
			fmt.Fprintln(out, 0)
			continue
		}
		if ww[h] == query(i, j) {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 0)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
