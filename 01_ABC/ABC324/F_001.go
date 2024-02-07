package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const N = 200005

type node struct {
	x    int
	y, z float64
}

var n, m int
var e [N][]node
var dp [N]float64

func f(g float64) bool {
	for i := 2; i <= n; i++ {
		dp[i] = -1e20
		for _, j := range e[i] {
			dp[i] = math.Max(dp[i], dp[j.x]+j.y-g*j.z)
		}
	}
	return dp[n] >= 0
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var u, v int
		var p, q float64
		fmt.Fscan(in, &u, &v, &p, &q)
		e[v] = append(e[v], node{u, p, q})
	}
	l := 0.0
	r := 2e9
	ans := 0.0
	for r-l > 1e-10 {
		mid := (l + r) / 2
		if !f(mid) {
			r = mid
		} else {
			ans = mid
			l = mid
		}
	}
	fmt.Println(ans)
}
