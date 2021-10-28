package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

var (
	N    int
	res  int
	seen []bool
	G    [][]int
)

func DFS(v int) int {
	res = (res + powMod(2, N-1) - 1 + mod) % mod
	sum := 1
	seen[v] = true
	for _, u := range G[v] {
		if !seen[u] {
			d := DFS(u)
			res = (res - powMod(2, d) + 1 + mod) % mod
			sum += d
		}
	}
	res = (res - powMod(2, N-sum) + 1 + mod) % mod
	return sum
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N)

	G = make([][]int, N)
	for i := 0; i < N-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	seen = make([]bool, N)
	DFS(0)
	fmt.Println(res * powMod(powMod(2, N), mod-2) % mod)
}
