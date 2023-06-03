package main

import (
	"bufio"
	"fmt"
	"os"
)

var G [100001][]int
var visited [100001]bool

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	for i := 0; i < M; i++ {
		var A, B int
		fmt.Fscan(in, &A, &B)
		G[A] = append(G[A], B)
		G[B] = append(G[B], A)
	}
	dfs(N)
	fmt.Println(N)
}

func dfs(x int) bool {
	visited[x] = true
	for _, i := range G[x] {
		if visited[i] == false {
			if i == 1 || dfs(i) {
				fmt.Printf("%d ", i)
				return true
			}
		}
	}
	return false
}
