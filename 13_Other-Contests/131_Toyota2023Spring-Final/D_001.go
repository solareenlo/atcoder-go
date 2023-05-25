package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscan(in, &a, &b)

	b--
	c := b / a
	d := b % a

	var x [400010]int
	for i := 0; i < a; i++ {
		fmt.Fscan(in, &x[i])
	}
	for i := 0; i < a; i++ {
		x[i+a] = x[i]
	}
	var y [600010]int
	for i := 0; i+1 < 2*a; i++ {
		y[i] = (x[i+1] - x[i] + a) % a
	}

	V := make([]int, 0)
	for i := 0; i+1 < 2*a; i++ {
		V = append(V, y[i])
	}
	V = append(V, 1234567)

	SA := SuffixArrayInt(V)

	W := make([]int, 0)
	start := 0
	for i := 0; i < len(SA); i++ {
		ind := SA[i]
		if 0 <= ind && ind < a {
			W = append(W, ind)
		}
	}

	d++
	prev := c
	L := 0
	R := len(W) - 1
	for i := 0; i+1 < a; i++ {
		min := L
		max := R
		cut := R + 1
		for min <= max {
			h := (min + max) / 2
			val := (prev + y[W[h]+i])
			if val >= a {
				cut = h
				max = h - 1
			} else {
				min = h + 1
			}
		}

		s2 := (R - cut + 1)
		if s2 >= d {
			L = cut
		} else {
			R = cut - 1
			d -= s2
		}
		next := prev + y[W[d+L-1]+i]

		min = L
		max = R
		l := L
		for min <= max {
			h := (min + max) / 2
			val := (prev + y[W[h]+i])
			if val >= next {
				l = h
				max = h - 1
			} else {
				min = h + 1
			}
		}
		min = L
		max = R
		r := R
		for min <= max {
			h := (min + max) / 2
			val := (prev + y[W[h]+i])
			if val <= next {
				r = h
				min = h + 1
			} else {
				max = h - 1
			}
		}
		d -= (l - L)
		L = l
		R = r
		prev = next % a
	}
	start = W[L]

	gap := (c - x[start] + a) % a
	for i := start; i < start+a; i++ {
		fmt.Printf("%d ", (x[i]+gap)%a)
	}
}

func SuffixArrayInt(s []int) []int {
	n := len(s)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(l, r int) bool {
		return s[idx[l]] < s[idx[r]]
	})
	s2 := make([]int, n)
	now := 0
	for i := 0; i < n; i++ {
		if i != 0 && s[idx[i-1]] != s[idx[i]] {
			now++
		}
		s2[idx[i]] = now
	}
	return SAIS(s2, now, 10, 40)
}

func SANaive(ss []int) []int {
	s := make([]int, len(ss))
	copy(s, ss)
	n := len(s)
	sa := make([]int, n)
	for i := 0; i < n; i++ {
		sa[i] = i
	}
	sort.Slice(sa, func(i, j int) bool {
		l, r := sa[i], sa[j]
		if l == r {
			return false
		}
		for l < n && r < n {
			if s[l] != s[r] {
				return s[l] < s[r]
			}
			l++
			r++
		}
		return l == n
	})
	return sa
}

func SADoubling(ss []int) []int {
	s := make([]int, len(ss))
	copy(s, ss)
	n := len(s)
	sa, rnk, tmp := make([]int, n), s, make([]int, n)
	for i := 0; i < n; i++ {
		sa[i] = i
	}
	for k := 1; k < n; k *= 2 {
		cmp := func(i, j int) bool {
			x, y := sa[i], sa[j]
			if rnk[x] != rnk[y] {
				return rnk[x] < rnk[y]
			}
			rx, ry := -1, -1
			if x+k < n {
				rx = rnk[x+k]
			}
			if y+k < n {
				ry = rnk[y+k]
			}
			return rx < ry
		}
		sort.Slice(sa, cmp)
		tmp[sa[0]] = 0
		for i := 1; i < n; i++ {
			tmp[sa[i]] = tmp[sa[i-1]]
			if cmp(i-1, i) {
				tmp[sa[i]]++
			}
		}
		tmp, rnk = rnk, tmp
	}
	return sa
}

func SAIS(ss []int, upper, thresholdNaive, thresholdDoubling int) []int {
	s := make([]int, len(ss))
	copy(s, ss)
	n := len(s)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		if s[0] < s[1] {
			return []int{0, 1}
		} else {
			return []int{1, 0}
		}
	}
	if n < thresholdNaive {
		return SANaive(s)
	}
	if n < thresholdDoubling {
		return SADoubling(s)
	}

	sa := make([]int, n)
	ls := make([]bool, n)
	for i := n - 2; i >= 0; i-- {
		if s[i] == s[i+1] {
			ls[i] = ls[i+1]
		} else {
			ls[i] = s[i] < s[i+1]
		}
	}
	suml, sums := make([]int, upper+1), make([]int, upper+1)
	for i := 0; i < n; i++ {
		if !ls[i] {
			sums[s[i]]++
		} else {
			suml[s[i]+1]++
		}
	}
	for i := 0; i <= upper; i++ {
		sums[i] += suml[i]
		if i < upper {
			suml[i+1] += sums[i]
		}
	}

	induce := func(lms []int) {
		for i := 0; i < len(sa); i++ {
			sa[i] = -1
		}
		buf := make([]int, upper+1)
		copy(buf, sums)
		for _, d := range lms {
			if d == n {
				continue
			}
			sa[buf[s[d]]] = d
			buf[s[d]]++
		}
		copy(buf, suml)
		sa[buf[s[n-1]]] = n - 1
		buf[s[n-1]]++
		for i := 0; i < n; i++ {
			v := sa[i]
			if v >= 1 && !ls[v-1] {
				sa[buf[s[v-1]]] = v - 1
				buf[s[v-1]]++
			}
		}
		copy(buf, suml)
		for i := n - 1; i >= 0; i-- {
			v := sa[i]
			if v >= 1 && ls[v-1] {
				sa[buf[s[v-1]+1]-1] = v - 1
				buf[s[v-1]+1]--
			}
		}
	}
	lmsMap := make([]int, n+1)
	for i := 0; i < len(lmsMap); i++ {
		lmsMap[i] = -1
	}
	m := 0
	for i := 1; i < n; i++ {
		if !ls[i-1] && ls[i] {
			lmsMap[i] = m
			m++
		}
	}
	lms := make([]int, 0, m)
	for i := 1; i < n; i++ {
		if !ls[i-1] && ls[i] {
			lms = append(lms, i)
		}
	}
	induce(lms)
	if m > 0 {
		sortedLms := make([]int, 0, m)
		for _, v := range sa {
			if lmsMap[v] != -1 {
				sortedLms = append(sortedLms, v)
			}
		}
		recs := make([]int, m)
		recUpper := 0
		recs[lmsMap[sortedLms[0]]] = 0
		for i := 1; i < m; i++ {
			l, r := sortedLms[i-1], sortedLms[i]
			endl, endr := n, n
			if lmsMap[l]+1 < m {
				endl = lms[lmsMap[l]+1]
			}
			if lmsMap[r]+1 < m {
				endr = lms[lmsMap[r]+1]
			}
			same := true
			if endl-l != endr-r {
				same = false
			} else {
				for l < endl {
					if s[l] != s[r] {
						break
					}
					l++
					r++
				}
				if l == n || s[l] != s[r] {
					same = false
				}
			}
			if !same {
				recUpper++
			}
			recs[lmsMap[sortedLms[i]]] = recUpper
		}
		recSa := SAIS(recs, recUpper, thresholdNaive, thresholdDoubling)
		for i := 0; i < m; i++ {
			sortedLms[i] = lms[recSa[i]]
		}
		induce(sortedLms)
	}
	return sa
}
