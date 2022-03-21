package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 200001

var (
	K   int
	cov = make([]int, N<<2)
	t   = make([]int, N)
	F   = make([]int, N)
	L   = make([]int, N)
)

func pushdown(x int) {
	if x != 0 && cov[x] != 0 {
		cov[x<<1] = cov[x]
		cov[x<<1|1] = cov[x]
		cov[x] = 0
	}
}

func Cov(x, l, r, ql, qr, v int) {
	if ql > qr {
		return
	}
	if l >= ql && r <= qr {
		cov[x] = v
		return
	}
	mid := (l + r) >> 1
	pushdown(x)
	if ql <= mid {
		Cov(x<<1, l, mid, ql, qr, v)
	}
	if qr > mid {
		Cov(x<<1|1, mid+1, r, ql, qr, v)
	}
}

func Q(x, l, r, v int) int {
	if l == r {
		return cov[x]
	}
	mid := (l + r) >> 1
	pushdown(x)
	if v <= mid {
		return Q(x<<1, l, mid, v)
	}
	return Q(x<<1|1, mid+1, r, v)
}

func q(x int) int {
	pos := Q(1, 1, t[0], x)
	if pos == 0 {
		return 0
	}
	return F[pos] + ((t[L[pos]]-t[x])%K+K)%K
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n, &K)

	a := make([]int, n+1)
	b := make([]int, n+1)
	A := make([]int, N)
	R := make([]int, N)
	tot := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		A[i] = a[i] + A[i-1]
		if b[i] == 1 {
			if 2*a[i] > K {
				fmt.Println(-1)
				return
			}
			t[0]++
			tot++
			L[tot] = ((K - 2*A[i-1]%K) + K) % K
			t[t[0]] = L[tot]
			t[0]++
			R[tot] = ((K - 2*A[i]%K) + K) % K
			t[t[0]] = R[tot]
		}
	}
	tmp := t[1 : t[0]+1]
	sort.Ints(tmp)

	t[0] = len(unique(tmp))
	for i := 1; i <= tot; i++ {
		L[i] = lowerBound(t[1:t[0]+1], L[i]) + 1
		R[i] = lowerBound(t[1:t[0]+1], R[i]) + 1
	}
	for i := tot; i > 0; i-- {
		F[i] = q(L[i])
		if L[i] > R[i] {
			Cov(1, 1, t[0], R[i]+1, L[i]-1, i)
		} else {
			Cov(1, 1, t[0], 1, L[i]-1, i)
			Cov(1, 1, t[0], R[i]+1, t[0], i)
		}
	}

	ans := 1 << 60
	for i := 1; i <= t[0]; i++ {
		ans = min(ans, q(i))
	}
	fmt.Println(ans + 2*A[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
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
