package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type P struct {
		x, y int
	}

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
	}
	res := 1 - k
	for _, p := range a {
		res += p
	}
	res *= 2
	res += 3 * n
	v := make([]P, 1)
	v[0] = P{0, 0}
	memo := make([]P, 0)
	for i := 0; i < n+1; i++ {
		nxt := i
		now := 0
		if i != n {
			now = a[i]
		}
		for v[len(v)-1].x > now {
			h := v[len(v)-1].x
			id := v[len(v)-1].y
			v = v[:len(v)-1]
			nxt = id
			memo = append(memo, P{i - id, h - max(v[len(v)-1].x, now)})
		}
		if v[len(v)-1].x < now {
			v = append(v, P{now, nxt})
		}
	}
	sort.Slice(memo, func(i, j int) bool {
		if memo[i].x == memo[j].x {
			return memo[i].y < memo[j].y
		}
		return memo[i].x < memo[j].x
	})

	Len := len(memo)
	for i := 0; i < Len; i++ {
		if memo[i].x != 0 {
			w := memo[i].x
			h := memo[i].y
			now := min(k/w, h)
			k -= now * w
			h -= now
			res += 2 * h
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
