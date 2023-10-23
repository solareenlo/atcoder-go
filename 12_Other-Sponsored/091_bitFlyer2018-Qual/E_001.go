package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var y, w, n, m, d int
	fmt.Fscan(in, &y, &w, &n, &m, &d)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
	}
	sort.Ints(a)

	b := make([]int, m+2)
	c := make([]int, m+2)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i], &c[i])
		b[i]--
		c[i]--
	}

	b[m] = -y * 10
	c[m] = 0
	b[m+1] = y * 10
	c[m+1] = 0

	vb := make([]int, 0)
	for i := 0; i < m+2; i++ {
		vb = append(vb, b[i])
	}
	sort.Ints(vb)
	vb = unique(vb)
	sz := len(vb)
	r := make(map[int]int)
	for i := 0; i < sz; i++ {
		r[vb[i]] = i
	}

	v := make([][]int, w+1)
	dp := make([][]int, sz)
	for i := 0; i < m+2; i++ {
		if i < m {
			v[c[i]] = append(v[c[i]], b[i])
		}
		dp[r[b[i]]] = append(dp[r[b[i]]], c[i])
	}

	for i := 0; i < sz; i++ {
		sort.Ints(dp[i])
	}

	ans := m

	var thr func(int) int
	thr = func(x int) int {
		if x <= d {
			return x
		}
		return 0
	}
	var calc func(int) int
	calc = func(x int) int {
		k := r[x]
		if k == 0 {
			return 0
		}
		dw := w * (vb[k] - vb[k-1] - 1)
		dw += w - dp[k-1][len(dp[k-1])-1]
		dw += dp[k][0] - 1
		return thr(dw)
	}

	for i := 0; i < sz; i++ {
		for j := 0; j+1 < len(dp[i]); j++ {
			ans += thr(dp[i][j+1] - dp[i][j] - 1)
		}
		ans += calc(vb[i])
	}

	ps := make([]int, n)
	for i := 0; i < n; i++ {
		ps[i] = lowerBound(vb, a[i]/w)
	}

	for t := 0; t < w; t++ {
		dif := 0
		for i := 0; i < n; i++ {
			p := a[i] / w
			q := a[i]%w + t
			k := ps[i]
			l := 0
			r := 0
			if vb[k] == p {
				if dp[k][len(dp[k])-1] >= q && dp[k][lowerBound(dp[k], q)] == q {
					continue
				}
				if dp[k][0] > q {
					l += (p - vb[k-1] - 1) * w
					l += w - dp[k-1][len(dp[k-1])-1]
					l += q - 1
				} else {
					l += q - dp[k][lowerBound(dp[k], q)-1] - 1
				}
				if dp[k][len(dp[k])-1] < q {
					r += (vb[k+1] - p - 1) * w
					r += w - q
					r += dp[k+1][0] - 1
				} else {
					r += dp[k][upperBound(dp[k], q)] - q - 1
				}
			} else {
				l += (p - vb[k-1] - 1) * w
				l += w - dp[k-1][len(dp[k-1])-1]
				l += q - 1
				r += (vb[k] - p - 1) * w
				r += w - q
				r += dp[k][0] - 1
			}
			if l+1+r <= d {
				continue
			}
			if i != 0 {
				if a[i]-a[i-1]-1 < l {
					l = 0
				}
			}
			if i+1 < n {
				r = min(r, a[i+1]-a[i]-1)
			}
			dif += thr(l) + 1 + thr(r)
		}

		fmt.Println(ans + dif)

		for _, x := range v[t] {
			ans -= calc(x)
		}
		for _, x := range v[t] {
			if dp[r[x]+1][0] != t {
				ans -= calc(vb[r[x]+1])
			}
		}

		for _, x := range v[t] {
			k := r[x]
			dp[k] = dp[k][1:]
			if len(dp[k]) != 0 {
				ans -= thr(dp[k][0] - t - 1)
				ans += thr(w + t - dp[k][len(dp[k])-1] - 1)
			}
			dp[k] = append(dp[k], w+t)
		}

		for _, x := range v[t] {
			ans += calc(x)
		}
		for _, x := range v[t] {
			if dp[r[x]+1][len(dp[r[x]+1])-1] != w+t {
				ans += calc(vb[r[x]+1])
			}
		}
	}
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
