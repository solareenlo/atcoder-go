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

	const INF = 1e77

	type car struct {
		x, v, d int
	}

	var n, L int
	fmt.Fscan(in, &n, &L)

	var vs, ds int
	fmt.Fscan(in, &vs, &ds)
	C := make([]car, 0)
	C = append(C, car{0, vs, ds})

	for i := 0; i < n; i++ {
		var x, v, d int
		fmt.Fscan(in, &x, &v, &d)
		C = append(C, car{x, v, d})
	}
	C = append(C, car{L, 0, 0})
	sort.Slice(C, func(i, j int) bool {
		return C[i].x < C[j].x
	})

	dp := make([]float64, n+2)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0.0
	for i := 0; i < n+1; i++ {
		if dp[i] < INF {
			for j := i + 1; j < n+2; j++ {
				if C[j].x <= C[i].x+C[i].d {
					dp[j] = math.Min(dp[j], dp[i]+float64(C[j].x-C[i].x)/float64(C[i].v))
				}
			}
		}
	}

	if dp[n+1] < INF {
		fmt.Println(dp[n+1])
	} else {
		fmt.Println("impossible")
	}
}
