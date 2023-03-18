package main

import (
	"fmt"
	"math"
	"sort"
)

type Pair struct {
	x, y int
}

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]Pair, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i].x, &p[i].y)
	}
	sort.Slice(p, func(i, j int) bool {
		return math.Atan2(float64(p[i].y), float64(p[i].x)) < math.Atan2(float64(p[j].y), float64(p[j].x))
	})
	ans := int(1e18)
	for len(p) != 0 {
		n = len(p)
		dp := make([]int, n)
		for i := range dp {
			dp[i] = int(1e18)
		}
		dp[0] = 0
		for i := 1; i < n; i++ {
			for j := i - 1; j >= 0; j-- {
				if dp[j] == int(1e18) {
					continue
				}
				t := int(p[j].x) * int(p[i].y)
				s := int(p[j].y) * int(p[i].x)
				if t-s <= 0 {
					continue
				}
				dp[i] = min(dp[i], dp[j]+abs(t-s))
			}
		}
		for j := 0; j < n; j++ {
			if dp[j] == int(1e18) {
				continue
			}
			i := 0
			t := int(p[j].x) * int(p[i].y)
			s := int(p[j].y) * int(p[i].x)
			if t-s <= 0 {
				continue
			}
			ans = min(ans, dp[j]+abs(t-s))
		}
		p = p[1:]
	}
	if ans == int(1e18) {
		fmt.Println("Impossible")
	} else {
		fmt.Println("Possible")
		pp := float64(ans) / 2.0
		fmt.Printf("%.10f\n", pp)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
