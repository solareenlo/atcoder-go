package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	var X, Y [17]float64
	for i := 0; i < N+M; i++ {
		fmt.Fscan(in, &X[i], &Y[i])
	}

	var dist [1 << 17][17]float64
	for i := 0; i < 1<<(N+M); i++ {
		for j := 0; j < N+M; j++ {
			dist[i][j] = 1e150
		}
	}

	for i := 0; i < N+M; i++ {
		dist[1<<i][i] = math.Hypot(X[i], Y[i])
	}

	ans := 1e150
	for i := 1; i < 1<<(N+M); i++ {
		for j := 0; j < N+M; j++ {
			if dist[i][j] < 1e100 {
				speed := float64(int(1) << bits.OnesCount(uint(i>>N)))
				if (i & ((1 << N) - 1)) == (1<<N)-1 {
					ans = math.Min(ans, dist[i][j]+math.Hypot(X[j], Y[j])/speed)
				}
				for k := 0; k < N+M; k++ {
					if ((i >> k) & 1) == 0 {
						dist[i|1<<k][k] = math.Min(dist[i|1<<k][k], dist[i][j]+math.Hypot(X[k]-X[j], Y[k]-Y[j])/(speed))
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
