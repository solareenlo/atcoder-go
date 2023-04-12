package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	edge := make([][]int, N)
	for M > 0 {
		M--
		var a, b int
		fmt.Fscan(in, &a, &b)
		edge[b-1] = append(edge[b-1], a-1)
	}
	for i := 0; i < N; i++ {
		edge[i] = append(edge[i], i)
	}
	tapi := make([]int, N)
	L, R := 0, N+1
	for i, j := 0, 0; i < N; i++ {
		for j < N {
			c := 0
			for _, k := range edge[j] {
				if tapi[k] == j-i {
					c++
				}
			}
			if c == 0 {
				break
			}
			for _, k := range edge[j] {
				tapi[k]++
			}
			j++
		}
		if j == N {
			break
		}
		if R-L > j-i {
			L = i
			R = j
		}
		for _, k := range edge[i] {
			tapi[k]--
		}
	}
	if R-L > N {
		fmt.Println(-1)
		return
	}
	fmt.Println(L+1, R+1)
}
