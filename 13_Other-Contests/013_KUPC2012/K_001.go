package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MA = 100005

type pair struct {
	x, y int
}

var G [][]pair
var rec []int
var seen []bool
var bas []int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M, Q int
	fmt.Fscan(in, &N, &M, &Q)
	G = make([][]pair, MA)
	rec = make([]int, MA)
	seen = make([]bool, MA)
	bas = make([]int, 0)
	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		G[a] = append(G[a], pair{b, c})
		G[b] = append(G[b], pair{a, c})
	}
	DFS(0)
	sort.Slice(bas, func(i, j int) bool {
		return bas[i] > bas[j]
	})
	for Q > 0 {
		Q--
		var a, b int
		fmt.Fscan(in, &a, &b)
		x := rec[a] ^ rec[b]
		for _, c := range bas {
			if x < (x ^ c) {
				x = x ^ c
			}
		}
		fmt.Fprintln(out, x)
	}
}

func DFS(u int) {
	seen[u] = true
	for _, to := range G[u] {
		if seen[to.x] {
			x := rec[u] ^ rec[to.x] ^ to.y
			add(x)
		} else {
			rec[to.x] = rec[u] ^ to.y
			DFS(to.x)
		}
	}
}

func add(x int) {
	for _, c := range bas {
		if x > (x ^ c) {
			x = x ^ c
		}
	}
	if x != 0 {
		bas = append(bas, x)
	}
}
