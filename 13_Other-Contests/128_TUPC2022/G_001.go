package main

import (
	"bufio"
	"fmt"
	"os"
)

var P []int

func main() {
	in := bufio.NewReader(os.Stdin)

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		P = make([]int, 0)
		var N int
		fmt.Fscan(in, &N)
		if N%2 == 1 {
			dfs(N, 0)
		} else {
			P = append(P, 0)
			P = append(P, 1)
			for u := 2; u < N; u += 2 {
				for v := 0; v < u; v++ {
					if v%2 == 0 {
						P = append(P, u)
					} else {
						P = append(P, u+1)
					}
					P = append(P, v)
				}
				P = append(P, u)
				P = append(P, u+1)
			}
		}
		for i := 0; i < len(P); i++ {
			fmt.Printf("%d ", P[i]+1)
		}
	}
}

func dfs(N, base int) {
	if N == 1 {
		P = append(P, base)
		return
	}
	P = append(P, base)
	P = append(P, base+1)
	P = append(P, base+2)
	P = append(P, base)
	for u := 3; u+1 < N; u += 2 {
		P = append(P, base+u)
		P = append(P, base+1)
		P = append(P, base+u+1)
		P = append(P, base+0)
	}
	P = P[:len(P)-1]
	dfs(N-2, base+2)
	P = P[:len(P)-1]
	P = append(P, base)
}
