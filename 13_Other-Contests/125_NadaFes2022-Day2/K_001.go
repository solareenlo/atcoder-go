package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const LINF = 1 << 62
	const IINF = 1 << 30

	var n, m, d int
	fmt.Fscan(in, &n, &m, &d)
	x := make([]int, n)
	v := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i])
	}

	max_num := 0
	px := -LINF
	for _, e := range x {
		if e-px >= d {
			px = e
			max_num++
		}
	}

	if max_num < m {
		fmt.Println("impossible")
		return
	}

	pre := make([]int, n)
	for l, r := 0, 0; r < n; r++ {
		for x[l]+d <= x[r] {
			l++
		}
		pre[r] = l
	}

	var f func(int) pair
	f = func(cost int) pair {
		dp := make([]pair, n+1)
		for i := range dp {
			dp[i] = pair{-LINF, -LINF}
		}
		dp[0] = pair{0, 0}
		for r := 0; r < n; r++ {
			dp[r+1] = maxPair(dp[r], plusPair(dp[pre[r]], pair{v[r] - cost, 1}))
		}
		tmp := dp[n]
		sum := tmp.x
		num := tmp.y
		return pair{sum + cost*num, num}
	}

	l := -IINF
	r := +IINF
	for r-l > 1 {
		cost := (l + r) >> 1
		tmp := f(cost)
		num := tmp.y
		if num >= m {
			l = cost
		} else {
			r = cost
		}
	}

	tmpL := f(l)
	sum_l := tmpL.x
	num_l := tmpL.y
	tmpR := f(r)
	sum_r := tmpR.x
	num_r := tmpR.y

	tan := (sum_l - sum_r) / (num_l - num_r)
	fmt.Println(sum_r + tan*(m-num_r))
}

type pair struct {
	x, y int
}

func plusPair(lhs, rhs pair) pair {
	return pair{lhs.x + rhs.x, lhs.y + rhs.y}
}

func maxPair(l, r pair) pair {
	if l.x == r.x {
		if l.y < r.y {
			return r
		}
		return l
	}
	if l.x < r.x {
		return r
	}
	return l
}
