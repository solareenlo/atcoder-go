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

	type point struct{ x, y int }
	p := make([]point, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i].x, &p[i].y)
	}
	sort.Slice(p, func(i, j int) bool {
		return p[j].x*(p[i].y-1) > p[i].x*(p[j].y-1)
	})

	res := 1
	p0 := p[0]
	for i := 1; i < n; i++ {
		p := p[i]
		if (p0.y-1)*(p.x-1) >= p.y*p0.x {
			res++
			p0 = p
		}
	}
	fmt.Println(res)
}
