package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

const N = 200009

var t [N][18]int
var a [N]int

func build(n int) {
	for i := 1; i <= n; i++ {
		t[i][0] = a[i]
	}
	for k := 1; k < 18; k++ {
		for i := 1; i+(1<<k)-1 <= n; i++ {
			t[i][k] = min(t[i][k-1], t[i+(1<<(k-1))][k-1])
		}
	}
}

func query(l, r int) int {
	k := 31 - countLeadingZeros(r-l+1)
	return min(t[l][k], t[r-(1<<k)+1][k])
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}
	sa := SuffixArrayInt(p)
	for i := 0; i < n; i++ {
		a[sa[i]+1] = i + 1
	}
	build(n)
	for i := 1; i <= n; i++ {
		l, r, ans := i, n, i
		for l <= r {
			mid := (l + r) >> 1
			if query(i, mid) >= a[i] {
				ans = mid
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countLeadingZeros(x int) int {
	return bits.LeadingZeros32(uint32(x))
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
