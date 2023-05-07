package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(2e18)

	var n, k, p int
	fmt.Fscan(in, &n, &k, &p)

	a := make([]int, n+1)
	for i := range a {
		a[i] = INF
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	n = (n + 1) / 2
	s := make([]pair, 1<<n)
	t := make([]pair, 1<<n)
	for i := 0; i < 1<<n; i++ {
		x := s[i].x
		y := s[i].y
		for j := n - 1; j >= 0 && i < (1<<j); j-- {
			s[i|(1<<j)] = pair{x + 1, y + a[j]}
			t[i|(1<<j)] = pair{x + 1, t[i].y + a[n+j]}
		}
	}
	sortPair(s)

	r := 0
	for _, tmp := range t {
		x := tmp.x
		y := tmp.y
		u := upperBound(s, pair{k - x, p - y})
		l := lowerBound(s, pair{k - x, 0})
		r += u - l
	}
	fmt.Println(r)
}

func lowerBound(a []pair, x pair) int {
	idx := sort.Search(len(a), func(i int) bool {
		if a[i].x == x.x {
			return a[i].y >= x.y
		}
		return a[i].x >= x.x
	})
	return idx
}

func upperBound(a []pair, x pair) int {
	idx := sort.Search(len(a), func(i int) bool {
		if a[i].x == x.x {
			return a[i].y > x.y
		}
		return a[i].x > x.x
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
