package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	type pair struct{ c, x int }
	p := make([]pair, 0)
	for x := 0; x < (1<<n)-1; x++ {
		var c int
		fmt.Fscan(in, &c)
		p = append(p, pair{c, x + 1})
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].c < p[j].c
	})

	ans := 0
	bases := make([]int, 0)
	for i := range p {
		x := p[i].x
		c := p[i].c
		for _, b := range bases {
			x = min(x, x^b)
		}
		if x != 0 {
			bases = append(bases, x)
			ans += c
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
