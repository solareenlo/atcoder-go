package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var N, M, K int
	fmt.Fscan(in, &N, &M, &K)
	P := make([][]pair, 301)
	P[0] = append(P[0], pair{0, N})
	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(in, &A)
		for j := K - 1; j >= 0; j-- {
			if len(P[j]) == 0 {
				continue
			}
			for P[j][0].y+M < i {
				P[j] = P[j][1:]
			}
			nxt := P[j][0].x + A*(j+1)
			for len(P[j+1]) != 0 && P[j+1][len(P[j+1])-1].x <= nxt {
				P[j+1] = P[j+1][:len(P[j+1])-1]
			}
			P[j+1] = append(P[j+1], pair{nxt, i})
		}
	}
	fmt.Println(P[K][0].x)
}
