package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

var E [400001]pair
var H, P [200001]int
var D [200001]bool

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	for i := 1; i < 2*M; i += 2 {
		var u, v int
		fmt.Fscan(in, &u, &v)
		E[i] = pair{v, H[u]}
		H[u] = i
		E[i+1] = pair{u, H[v]}
		H[v] = i + 1
	}
	dfs(1)

	if P[1] == 0 {
		p := E[H[1]].x
		P[1] = P[p]
		P[p] = 1
	}
	for i := 1; i < N+1; i++ {
		fmt.Printf("%d ", P[i])
	}
}

func dfs(A int) int {
	if !D[A] {
		D[A] = true
		k := 0
		for i := H[A]; i > 0; i = E[i].y {
			if k == 0 {
				k = dfs(E[i].x)
			} else {
				t := dfs(E[i].x)
				if t != 0 {
					P[k] = t
					P[t] = k
					k = 0
				}
			}
		}
		if k != 0 {
			P[k] = A
			P[A] = k
			return 0
		} else {
			return A
		}
	}
	return 0
}
