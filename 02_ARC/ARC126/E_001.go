package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 5000005
const mod = 998244353

var (
	size = [N]int{}
	sum  = [N]int{}
	sall = [N]int{}
	b    = [600600]int{}
)

func Add(p, l, r, x, y int) {
	if l == r {
		size[p] += y
		sum[p] = size[p] * b[l] % mod
		sall[p] = b[l] * size[p] % mod * (size[p] + 1) % mod * ((mod + 1) / 2) % mod
		return
	}
	mid := (l + r) / 2
	if x <= mid {
		Add(p*2, l, mid, x, y)
	} else {
		Add(p*2+1, mid+1, r, x, y)
	}
	sum[p] = (sum[p*2] + sum[p*2+1]) % mod
	sall[p] = (sall[p*2] + sall[p*2+1] + size[p*2]*sum[p*2+1]) % mod
	size[p] = size[p*2] + size[p*2+1]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)
	tmp := (n + 1) * (mod + 1) / 2 % mod

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		b[0]++
		b[b[0]] = a[i]
	}

	X := make([]int, q+1)
	Y := make([]int, q+1)
	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &X[i], &Y[i])
		b[0]++
		b[b[0]] = Y[i]
	}
	c := b[1 : b[0]+1]
	sort.Ints(c)

	c = unique(b[1 : b[0]+1])
	for i := 0; i < len(c); i++ {
		b[i+1] = c[i]
	}
	b[0] = len(c)
	for i := 1; i <= n; i++ {
		a[i] = LowerBound(b[1:b[0]+1], a[i]) + 1
		Add(1, 1, b[0], a[i], 1)
	}
	for i := 1; i <= q; i++ {
		Y[i] = LowerBound(b[1:b[0]+1], Y[i]) + 1
		Add(1, 1, b[0], a[X[i]], -1)
		Add(1, 1, b[0], Y[i], 1)
		a[X[i]] = Y[i]
		fmt.Fprintln(out, (sall[1]-sum[1]*tmp%mod+mod)%mod)
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
	sort.Ints(result)
	return result
}

func LowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
