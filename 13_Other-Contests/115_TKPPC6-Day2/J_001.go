package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)
	const amax = 1000000

	var nxt [amax + 5]int
	for i := amax; i >= 2; i-- {
		for j := i; j <= amax; j += i {
			nxt[j] = i
		}
	}
	prime := make([]int, 0)
	for i := 2; i <= 1000; i++ {
		if nxt[i] == i {
			prime = append(prime, i)
		}
	}
	m := len(prime)
	var n, q int
	fmt.Fscan(in, &n, &q)
	var dist [170][170]int
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i != j {
				dist[i][j] = INF
			}
		}
	}
	var idx [1010]int
	for i := 0; i < m; i++ {
		idx[prime[i]] = i
	}
	var pf, big [amax + 5][]int
	var a [100010]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] == 1 || len(pf[a[i]]) != 0 {
			continue
		}
		x := a[i]
		for x > 1 {
			pf[a[i]] = append(pf[a[i]], nxt[x])
			x /= nxt[x]
		}
		x = a[i]
		pf[x] = unique(pf[x])
		if pf[x][len(pf[x])-1] > 1000 {
			for _, j := range pf[x] {
				tmp := pf[x][len(pf[x])-1]
				big[tmp] = append(big[tmp], j)
			}
			tmp := pf[x][len(pf[x])-1]
			big[tmp] = big[tmp][:len(big[tmp])-1]
		}
		for _, f := range pf[x] {
			if f > 1000 {
				continue
			}
			for _, g := range pf[x] {
				if g > 1000 {
					continue
				}
				dist[idx[f]][idx[g]] = min(dist[idx[f]][idx[g]], g)
			}
		}
	}
	for i := 1000; i <= amax; i++ {
		sort.Ints(big[i])
		unique(big[i])
		for _, f := range big[i] {
			for _, g := range big[i] {
				dist[idx[f]][idx[g]] = min(dist[idx[f]][idx[g]], i+g)
			}
		}
	}
	for k := 0; k < m; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	for q > 0 {
		q--
		var s, t int
		fmt.Fscan(in, &s, &t)
		s--
		t--
		x := a[s]
		y := a[t]
		if x == 1 || y == 1 {
			fmt.Println(-1)
			continue
		}
		ds := make([]int, m)
		dt := make([]int, m)
		for i := 0; i < m; i++ {
			ds[i] = INF
			dt[i] = INF
		}
		bs := pf[x][len(pf[x])-1]
		bt := pf[y][len(pf[y])-1]
		if bs > 1000 {
			for _, f := range big[bs] {
				ds[idx[f]] = min(ds[idx[f]], bs)
			}
		}
		for _, f := range pf[x] {
			if f > 1000 {
				continue
			}
			ds[idx[f]] = min(ds[idx[f]], 0)
		}
		if bt > 1000 {
			for _, g := range big[bt] {
				dt[idx[g]] = min(dt[idx[g]], bt)
			}
		}
		for _, g := range pf[y] {
			if g > 1000 {
				continue
			}
			dt[idx[g]] = min(dt[idx[g]], 0)
		}
		ans := INF
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				ans = min(ans, ds[i]+prime[i]+dist[i][j]+dt[j])
			}
		}
		if gcd(a[s], a[t]) != 1 {
			ans = min(ans, nxt[gcd(a[s], a[t])])
		}
		if ans == INF {
			fmt.Println(-1)
		} else {
			fmt.Println(ans)
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
