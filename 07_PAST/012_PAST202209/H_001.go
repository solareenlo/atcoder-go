package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscan(in, &n, &x)
	m := int(1e9)
	dp := make([]pair, x+1)
	dp[x] = pair{0, m}

	for n > 0 {
		n--
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		for i := b; i <= x; i++ {
			g, s := dp[i].x, dp[i].y
			dp[i-b] = maxPair(dp[i-b], pair{g + c, s - a})
		}
	}

	var ans tuple
	for i := 0; i < x+1; i++ {
		g, s := dp[i].x, dp[i].y
		ans = maxTuple(ans, tuple{g, s, i})
	}
	fmt.Println(ans.x, ans.y, ans.z)
}

type tuple struct {
	x, y, z int
}

func maxTuple(a, b tuple) tuple {
	if a.x == b.x {
		if a.y == b.y {
			if a.z > b.z {
				return a
			}
			return b
		}
		if a.y > b.y {
			return a
		}
		return b
	}
	if a.x > b.x {
		return a
	}
	return b
}

type pair struct {
	x, y int
}

func maxPair(a, b pair) pair {
	if a.x == b.x {
		if a.y > b.y {
			return a
		}
		return b
	}
	if a.x > b.x {
		return a
	}
	return b
}
