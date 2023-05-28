package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INFLL = 1 << 62
	const INF = 1 << 30

	var n int
	fmt.Fscan(in, &n)
	x := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i].x, &x[i].y)
	}
	ans := INFLL
	for j := 0; j < 2; j++ {
		sortPair(x)
		mi := INF
		ma := -INF
		for l := 0; l < n; l++ {
			mmi := mi
			mma := ma
			for r := n; l < r; r-- {
				ans = min(ans, f(x[l].x, x[r-1].x, mmi, mma))
				mmi = min(mmi, x[r-1].y)
				mma = max(mma, x[r-1].y)
			}
			mi = min(mi, x[l].y)
			ma = max(ma, x[l].y)
		}
		for i := 0; i < n; i++ {
			x[i].x, x[i].y = x[i].y, x[i].x
		}
	}
	fmt.Printf("%10.10f\n", math.Sqrt(float64(ans)))
}

func f(x1, x2, y1, y2 int) int {
	if x1 > x2 {
		x1 = 0
		x2 = 0
	}
	if y1 > y2 {
		y1 = 0
		y2 = 0
	}
	return d(x2-x1+min(abs(x1), abs(x2)), y2-y1+min(abs(y1), abs(y2)))
}

func d(x, y int) int { return x*x + y*y }

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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
