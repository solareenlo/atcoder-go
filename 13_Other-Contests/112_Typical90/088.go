package main

import (
	"bufio"
	"fmt"
	"os"
)

var N int
var A, out [88]int
var G [88][]int
var now []int
var sum int
var S [8889][]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var Q int
	fmt.Fscan(in, &N, &Q)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	for i := 0; i < Q; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}
	dfs(0)
}

func dfs(id int) {
	for k := id; k < N; k++ {
		if out[k] == 0 {
			for _, v := range G[k] {
				out[v]++
			}
			now = append(now, k)
			sum += A[k]
			if len(S[sum]) != 0 {
				disp(S[sum])
				disp(now)
				os.Exit(0)
			}
			S[sum] = make([]int, len(now))
			copy(S[sum], now)
			dfs(k + 1)
			now = now[:len(now)-1]
			sum -= A[k]
			for _, v := range G[k] {
				out[v]--
			}
		}
	}
}

func disp(B []int) {
	fmt.Println(len(B))
	for i := 0; i < len(B); i++ {
		if i+1 == len(B) {
			fmt.Println(B[i] + 1)
		} else {
			fmt.Printf("%d ", B[i]+1)
		}
	}
}
