package main

import (
	"bufio"
	"fmt"
	"os"
)

var G [1000][]int
var col [1000]int
var ok bool

func dfs(u, c int, a, b *int) {
	if c == 1 {
		(*a)++
	} else {
		(*b)++
	}
	col[u] = c
	for _, v := range G[u] {
		if col[v] == 0 {
			dfs(v, 3-c, a, b)
		} else if c == col[v] {
			ok = false
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var dp [340][340]bool

	var N, M int
	fmt.Fscan(in, &N, &M)
	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}
	C := [3]int{N / 3, (N + 1) / 3, (N + 2) / 3}
	dp[0][0] = true
	for i := 0; i < N; i++ {
		if col[i] == 0 {
			a, b := 0, 0
			ok = true
			dfs(i, 1, &a, &b)
			if ok {
				for j := C[1]; j >= 0; j-- {
					for k := C[2]; k >= 0; k-- {
						if dp[j][k] {
							if j+a <= C[1] && k+b <= C[2] {
								dp[j+a][k+b] = true
							}
							if j+b <= C[1] && k+a <= C[2] {
								dp[j+b][k+a] = true
							}
						}
					}
				}
			}
		}
	}
	if dp[C[1]][C[2]] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
