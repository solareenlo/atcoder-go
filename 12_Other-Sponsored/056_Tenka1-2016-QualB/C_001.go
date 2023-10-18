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

	type P struct {
		x, y int
	}

	var dp, C [50][50]float64
	C[0][0] = 1
	for i := 1; i < 50; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				C[i][j] = 1
			} else {
				C[i][j] = C[i-1][j-1] + C[i-1][j]
			}
		}
	}
	var n, p, q int
	fmt.Fscanf(in, "%d %d/%d", &n, &p, &q)
	v := make([]P, 0)
	var a [50]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var b int
			fmt.Fscan(in, &b)
			a[i] += b
		}
		v = append(v, P{a[i], i})
	}
	sort.Slice(v, func(a, b int) bool {
		if v[a].x != v[b].x {
			return v[a].x > v[b].x
		}
		return v[a].y < v[b].y
	})
	d := float64(p) / float64(q)
	dp[0][0] = 1.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 0 {
				continue
			}
			s := v[i].y
			for k := 0; k < a[s]+1; k++ {
				for t := 0; t < n-a[s]; t++ {
					if i != 0 {
						if (k+t > j) || (k+t == j && v[i-1].y > v[i].y) {
							continue
						}
					}
					dp[i+1][k+t] += dp[i][j] * C[a[s]][k] * math.Pow(d, float64(k)) * math.Pow(1-d, float64(a[s]-k)) * C[n-1-a[s]][t] * math.Pow(1-d, float64(t)) * math.Pow(d, float64(n-1-a[s]-t))
				}
			}
		}
	}
	ans := 0.0
	for i := 0; i < n; i++ {
		ans += dp[n][i]
	}
	fmt.Printf("%.12f\n", ans)
}
