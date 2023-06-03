package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200000

	var n, k int
	fmt.Fscan(in, &n, &k)
	l := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l[i], &r[i])
	}
	for i := 0; i < n; i++ {
		r[i] += k
	}

	p1 := make([]pair, n)
	for i := 0; i < n; i++ {
		p1[i] = pair{r[i], l[i]}
	}
	sortPair(p1)

	now := 0
	cnt := 0
	cntL := make([]int, N)
	for i := 0; i < n; i++ {
		nr := p1[i].x
		nl := p1[i].y
		if now <= nl {
			now = nr
			cnt++
			cntL[now] = cnt
		}
	}

	p2 := make([]pair, n)
	for i := 0; i < n; i++ {
		p2[i] = pair{l[i], r[i]}
	}
	sortPairR(p2)

	now = N
	cnt = 0
	cntR := make([]int, N)
	for i := 0; i < n; i++ {
		nl := p2[i].x
		nr := p2[i].y
		if now >= nr {
			now = nl
			cnt++
			cntR[now] = cnt
		}
	}

	for i := 1; i < N; i++ {
		cntL[i] = max(cntL[i], cntL[i-1])
	}
	for i := N - 2; i >= 0; i-- {
		cntR[i] = max(cntR[i], cntR[i+1])
	}

	for i := 0; i < n; i++ {
		fmt.Println(cntL[l[i]] + 1 + cntR[r[i]])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func sortPairR(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y > tmp[j].y
		}
		return tmp[i].x > tmp[j].x
	})
}
