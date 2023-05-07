package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	v := make([]tuple, N)
	for i := 0; i < N; i++ {
		var d, c, s int
		fmt.Fscan(in, &d, &c, &s)
		v[i] = tuple{d, c, s}
	}
	sortTuple(v)

	var dp [5002]int
	for i := 0; i < N; i++ {
		d := v[i].x
		c := v[i].y
		s := v[i].z
		for i := d - c; i >= 0; i-- {
			if dp[i+c] < dp[i]+s {
				dp[i+c] = dp[i] + s
			}
		}
	}
	res := 0
	for i := range dp {
		res = max(res, dp[i])
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type tuple struct {
	x, y, z int
}

func sortTuple(tup []tuple) {
	sort.Slice(tup, func(i, j int) bool {
		if tup[i].x == tup[j].x {
			if tup[i].y == tup[j].y {
				return tup[i].z < tup[j].z
			}
			return tup[i].y < tup[j].y
		}
		return tup[i].x < tup[j].x
	})
}
