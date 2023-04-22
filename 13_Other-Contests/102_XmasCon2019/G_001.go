package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(in, &H, &W)
	const LEFT = 29 * 29
	const N = LEFT + 29*29
	var m BipartiteMatching
	m.init(N)
	for y := 0; y < H; y++ {
		var s string
		fmt.Fscan(in, &s)
		for x := 0; x < W; x++ {
			if s[x] != '.' {
				continue
			}
			x1 := 2*x + 5*y
			y1 := 5*x - 2*y
			y1 %= 29
			if y1 < 0 {
				y1 += 29
			}
			b1 := (x1%29)*29 + (y1 % 29)
			x2 := 2*x - 5*y
			y2 := 5*x + 2*y
			x2 %= 29
			if x2 < 0 {
				x2 += 29
			}
			b2 := (x2%29)*29 + (y2 % 29)
			m.add_edge(b1, LEFT+b2)
		}
	}
	fmt.Println(m.matching())
}

type BipartiteMatching struct {
	G     [][]int
	match []int
	used  []bool
}

func (b *BipartiteMatching) init(N int) {
	b.G = make([][]int, N)
}

func (b *BipartiteMatching) add_edge(u, v int) {
	b.G[u] = append(b.G[u], v)
	b.G[v] = append(b.G[v], u)
}

func (b *BipartiteMatching) dfs(u int) bool {
	b.used[u] = true
	for _, v := range b.G[u] {
		w := b.match[v]
		if w < 0 || (!b.used[w] && b.dfs(w)) {
			b.match[u] = v
			b.match[v] = u
			return true
		}
	}
	return false
}

func (b *BipartiteMatching) matching() int {
	res := 0
	b.match = make([]int, len(b.G))
	for i := range b.match {
		b.match[i] = -1
	}
	for u := 0; u < len(b.G); u++ {
		if b.match[u] < 0 {
			b.used = make([]bool, len(b.G))
			if b.dfs(u) {
				res++
			}
		}
	}
	return res
}
