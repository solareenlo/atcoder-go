package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func inverse(x int) int {
	a := 1
	a = 2*a - a*a*x
	a = 2*a - a*a*x
	a = 2*a - a*a*x
	a = 2*a - a*a*x
	a = 2*a - a*a*x
	return a
}

const N = 200000 * 60

var trie [N][2]int
var sum [N]int

func dfs(x int) int {
	if x == 0 {
		return 0
	}
	return sum[x] + max(dfs(trie[x][0]), dfs(trie[x][1]))
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, L int
	fmt.Fscan(in, &N, &L)

	A := make([]int, N)
	V := make([]int, N)
	W := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i], &V[i], &W[i])
		A[i] = (1 << L) - A[i]
		A[i] &= (1 << L) - 1
		A[i] <<= L
	}

	cnt := 1
	for i := 0; i < N; i++ {
		tv := ctz(V[i])
		x := inverse(V[i]>>tv) * (A[i] >> tv)
		p := 1
		for i := 0; i < 2*L-tv; i++ {
			d := (x >> i) & 1
			if trie[p][d] == 0 {
				cnt++
				trie[p][d] = cnt
			}
			p = trie[p][d]
		}
		sum[p] += W[i]
	}

	fmt.Println(dfs(1))
}

func ctz(x int) int {
	return bits.TrailingZeros64(uint64(x))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
