package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 200003

var tot int
var nxt, ver [N << 1]int
var hd [N]int

func add(u, v int) {
	tot++
	nxt[tot] = hd[u]
	hd[u] = tot
	ver[tot] = v
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var s, t [N]bool
	var dis, que [N]int

	var n int
	fmt.Fscan(in, &n)
	var pre int
	fmt.Fscan(in, &pre)
	for i := 1; i < n; i++ {
		var cur int
		fmt.Fscan(in, &cur)
		add(pre, cur)
		add(cur, pre)
		pre = cur
	}
	var na int
	fmt.Fscan(in, &na)
	str := make([]int, na+1)
	for i := 0; i < na; i++ {
		fmt.Fscan(in, &str[i])
		s[str[i]] = true
	}
	str[na] = 0
	var nb int
	fmt.Fscan(in, &nb)
	resize(&str, na+nb+1)
	for i := 0; i < nb; i++ {
		fmt.Fscan(in, &str[na+1+i])
		t[str[na+1+i]] = true
	}
	sa := SuffixArrayInt(str)
	lcp := LcpArrayInt(str, sa)
	Len := 0
	for i := 0; i < na+nb; i++ {
		if (sa[i] < na && sa[i+1] > na) || (sa[i] > na && sa[i+1] < na) {
			Len = max(Len, lcp[i])
		}
	}
	if Len != 0 {
		fmt.Println(na + nb - Len - Len)
		return
	}
	for i := 1; i <= n; i++ {
		dis[i] = INF
	}
	tl := 0
	for i := 1; i <= n; i++ {
		if s[i] {
			dis[i] = 0
			tl++
			que[tl] = i
		}
	}
	for pos := 1; pos <= tl; pos++ {
		u := que[pos]
		for i := hd[u]; i > 0; i = nxt[i] {
			v := ver[i]
			if dis[v] < INF {
				continue
			}
			dis[v] = dis[u] + 1
			tl++
			que[tl] = v
		}
	}
	res := INF
	for i := 1; i <= n; i++ {
		if t[i] {
			res = min(res, dis[i])
		}
	}
	fmt.Println(na + nb + res + res - 2)
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

func SuffixArrayString(s string) []int {
	n := len(s)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		s2[i] = int(s[i])
	}
	return SAIS(s2, 255, 10, 40)
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

func LcpArrayInt(s, sa []int) []int {
	n := len(s)
	if n < 1 {
		panic("length of slice s must be more than or equal to 1")
	}
	rnk := make([]int, n)
	for i := 0; i < n; i++ {
		rnk[sa[i]] = i
	}
	lcp := make([]int, n-1)
	h := 0
	for i := 0; i < n; i++ {
		if h > 0 {
			h--
		}
		if rnk[i] == 0 {
			continue
		}
		j := sa[rnk[i]-1]
		for ; j+h < n && i+h < n; h++ {
			if s[j+h] != s[i+h] {
				break
			}
		}
		lcp[rnk[i]-1] = h
	}
	return lcp
}

func LcpArrayString(s string, sa []int) []int {
	n := len(s)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		s2[i] = int(s[i])
	}
	return LcpArrayInt(s2, sa)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
