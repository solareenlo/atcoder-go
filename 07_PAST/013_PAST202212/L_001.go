package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e9) + 7

	var n int
	fmt.Fscan(in, &n)
	P := make([]pair, n)
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		P[i] = pair{r, -l}
	}
	sortPair(P)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = P[i].y
	}
	L := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		L[i] = INF
	}
	for i := 0; i < n; i++ {
		idx := upperBound(L, A[i])
		L[idx] = A[i]
	}
	fmt.Println(lowerBound(L, INF))
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

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}
