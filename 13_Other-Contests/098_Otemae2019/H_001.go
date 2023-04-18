package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const mod = 998244353
const INF = 2000000000000005

var N int
var A [200005]int
var dpl, dpr [2][200005]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	fmt.Println(dfs(0, N))
}

func dfs(l, r int) int {
	if r-l == 1 {
		return (A[l]) % mod
	}
	res := (dfs(l, (l+r)/2) + dfs((l+r)/2, r)) % mod
	dpl[0][((l+r)/2)-1] = 0
	dpl[0][(l+r)/2] = 0
	dpl[1][((l+r)/2)-1] = A[((l+r)/2)-1]
	dpl[1][((l + r) / 2)] = -INF
	dpr[0][((l+r)/2)-1] = 0
	dpr[0][(l+r)/2] = 0
	dpr[1][((l+r)/2)-1] = -INF
	dpr[1][((l + r) / 2)] = A[((l + r) / 2)]
	a := make([]int, 0)
	b := make([]int, 0)
	a = append(a, A[((l+r)/2)-1])
	b = append(b, A[(l+r)/2])
	for i := ((l + r) / 2) - 2; l <= i; i-- {
		dpl[0][i] = max(dpl[0][i+1], A[i]+dpl[0][i+2])
		dpl[1][i] = max(dpl[1][i+1], A[i]+dpl[1][i+2])
		res = (res + ((dpl[0][i])%mod)*(r-(l+r)/2)) % mod
		if dpl[0][i] < dpl[1][i] {
			a = append(a, dpl[1][i]-dpl[0][i])
		}
	}
	for i := ((l + r) / 2) + 1; i < r; i++ {
		dpr[0][i] = max(dpr[0][i-1], A[i]+dpr[0][i-2])
		dpr[1][i] = max(dpr[1][i-1], A[i]+dpr[1][i-2])
		res = (res + ((dpr[0][i])%mod)*(((l+r)/2)-l)) % mod
		if dpr[0][i] < dpr[1][i] {
			b = append(b, dpr[1][i]-dpr[0][i])
		}
	}
	n := len(a)
	m := len(b)
	sort.Ints(a)
	sort.Ints(b)
	for i := 0; i < n; i++ {
		k := UpperBound(b, a[i])
		k = (r - (l+r)/2 - m + k)
		res = (res + ((a[i])%mod)*k) % mod
	}
	for i := 0; i < m; i++ {
		k := LowerBound(a, b[i])
		k = (((l + r) / 2) - l - n + k)
		res = (res + ((b[i])%mod)*k) % mod
	}
	return res
}

func UpperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func LowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
