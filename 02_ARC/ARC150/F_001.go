package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MAXM = 200005

var app [MAXM][]int
var N int
var dp, tag [MAXM]int

func FindSucc(x, v int) int {
	p := (x - 1) / N
	r := (x-1)%N + 1
	if r >= app[v][len(app[v])-1] {
		return app[v][0] + (p+1)*N
	}
	idx := upperBound(app[v], r)
	return p*N + app[v][idx]
}

func FindPred(x, v int) int {
	p := (x - 1) / N
	r := (x-1)%N + 1
	if r < app[v][0] {
		return app[v][len(app[v])-1] + (p-1)*N
	}
	idx := upperBound(app[v], r) - 1
	return p*N + app[v][idx]
}

func Divide(l, r int) {
	if l == r {
		return
	}
	mid := (l + r) >> 1
	Divide(l, mid)
	for i := mid; i <= r; i++ {
		tag[i] = 0
	}
	for i := 1; i <= r-l; i++ {
		rig := min(r-i, mid)
		tmp := FindPred(dp[rig], i)
		lef := lowerBound(dp[l:mid+1], tmp) + l
		tag[lef+i] = max(tag[lef+i], FindSucc(dp[rig], i))
	}
	for i := mid + 1; i <= r; i++ {
		tag[i] = max(tag[i], tag[i-1])
		dp[i] = max(dp[i], tag[i])
	}
	Divide(mid+1, r)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var M int
	fmt.Fscan(in, &N, &M)
	for i := 1; i <= N; i++ {
		var a int
		fmt.Fscan(in, &a)
		app[a] = append(app[a], i)
	}
	dp[0] = 0
	Divide(0, M)
	fmt.Println(dp[M])
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func lowerBound(a []int, x int) int {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
