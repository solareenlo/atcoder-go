package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &h[i])
	}

	v := make([]tuple, 0)
	Var := make([][]tuple, 3010)

	for i := 0; i < n; i++ {
		var m, s, e int
		fmt.Fscan(in, &m, &s, &e)
		m--
		v = append(v, tuple{s, e, m})
	}
	sortTuple(v)

	for i := 0; i < n; i++ {
		s := v[i].x
		e := v[i].y
		m := v[i].z
		Var[m] = append(Var[m], tuple{e, s, i + 1})
	}
	for i := 0; i < n; i++ {
		sortTuple(Var[i])
	}

	var dp [3010]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			en := 0
			if i != 0 {
				en = v[i-1].y
			}
			sum := 0
			cnt := 0
			for _, e := range Var[j] {
				ee := e.x
				ss := e.y
				ii := e.z
				if en <= ss {
					sum += h[cnt]
					dp[ii] = max(dp[ii], dp[i]+sum)
					en = ee
					cnt++
				}
			}
		}
	}

	ans := 0
	for i := 0; i < n+1; i++ {
		ans = max(ans, dp[i])
	}
	fmt.Println(ans)
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
