package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const mod = 1000000007

var N, K, M int
var inv [7]int
var a [4]int
var xs []int
var can [16][20]bool
var asn [5][4][]int
var bt, sel [4]int
var ans int

func main() {
	in := bufio.NewReader(os.Stdin)

	for i := 0; i < 7; i++ {
		a := 1
		for j := 0; j < i; j++ {
			if a%i == 0 {
				inv[i] = a / i
				break
			}
			a += mod
		}
	}

	fmt.Fscan(in, &N, &K)
	for i := 0; i < K; i++ {
		fmt.Fscan(in, &a[i])
	}
	xs = make([]int, 0)
	for i := 0; i < K; i++ {
		xs = append(xs, a[i])
	}
	for i := 0; i < 1<<K; i++ {
		if i == 0 {
			continue
		}
		s := 0
		for j := 0; j < K; j++ {
			if B(i, j) {
				s += a[j]
			}
		}
		xs = append(xs, s+1)
	}
	sort.Ints(xs)
	xs = unique(xs)
	M = len(xs)
	for i := 0; i < 1<<K; i++ {
		if i == 0 {
			continue
		}
		l, r := 0, 0
		for j := 0; j < K; j++ {
			if B(i, j) {
				l = max(l, a[j])
				r += a[j]
			}
		}
		for j := 0; j < M; j++ {
			if l <= xs[j] && xs[j] <= r {
				can[i][j] = true
			}
		}
	}
	for i := range asn {
		for j := range asn[i] {
			asn[i][j] = make([]int, 0)
		}
	}
	dfsas(0, (1<<K)-1)
	dfs(0)
	if ans < 0 {
		ans += mod
	}
	fmt.Println(ans)
}

func dfsas(t, left int) {
	if left == 0 {
		for i := 0; i < t; i++ {
			asn[t][i] = append(asn[t][i], bt[i])
		}
	}
	if t == K {
		return
	}
	for i := 0; i < 1<<K; i++ {
		if i > 0 && (i&left) == i {
			bt[t] = i
			dfsas(t+1, left^i)
		}
	}
}

func dfs(t int) {
	if t > 0 {
		ok := false
		A := len(asn[t][0])
		for i := 0; i < A; i++ {
			no := false
			for j := 0; j < t; j++ {
				if !can[asn[t][j][i]][sel[j]] {
					no = true
				}
			}
			if !no {
				ok = true
			}
		}
		if ok {
			L := N + 1
			vc := make([]int, 0)
			for j := 0; j < t; j++ {
				L -= xs[sel[j]]
				if xs[sel[j]]+1 != xs[sel[j]+1] {
					vc = append(vc, xs[sel[j]+1]-xs[sel[j]])
				}
			}
			tmp := 0
			if len(vc) == 0 {
				tmp = C(L, t)
			}
			if len(vc) == 1 {
				tmp = C(L+1, t+1) - C(L-vc[0]+1, t+1)
			}
			if len(vc) == 2 {
				tmp = C(L+2, t+2) - C(L-vc[0]+2, t+2) - C(L-vc[1]+2, t+2) + C(L-vc[0]-vc[1]+2, t+2)
			}
			ans = (ans + tmp) % mod
		}
	}
	if t == K {
		return
	}
	for i := 0; i < M; i++ {
		sel[t] = i
		dfs(t + 1)
	}
}

func B(x, y int) bool {
	if (x>>y)&1 != 0 {
		return true
	}
	return false
}

func C(x, y int) int {
	if x < y {
		return 0
	}
	a := 1
	for i := 0; i < y; i++ {
		a = a * (x - i) % mod * inv[i+1] % mod
	}
	return a
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
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
