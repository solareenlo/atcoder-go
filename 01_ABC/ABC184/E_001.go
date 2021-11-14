package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)

	G := [26][]pair{}
	c := make([][]byte, h+2)
	for i := range c {
		c[i] = make([]byte, w+2)
		for j := range c[i] {
			c[i][j] = '#'
		}
	}
	var start pair
	for i := 1; i <= h; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := range s {
			c[i][j+1] = s[j]
		}
		for j := 1; j <= w; j++ {
			if c[i][j] >= 'a' && c[i][j] <= 'z' {
				G[int(c[i][j]-'a')] = append(G[int(c[i][j]-'a')], pair{i, j})
			}
			if c[i][j] == 'S' {
				start = pair{i, j}
			}
		}
	}

	vis := make([][]bool, h+1)
	for i := range vis {
		vis[i] = make([]bool, w+1)
	}

	q := make([]pd, 0, h*w)
	q = append(q, pd{start, 0})
	for len(q) > 0 {
		x := q[0].p.x
		y := q[0].p.y
		d := q[0].d
		q = q[1:]
		ch := c[x][y]
		if x < 1 || y < 1 || x > h || y > w || vis[x][y] || ch == '#' {
			continue
		}
		vis[x][y] = true
		if ch == 'G' {
			fmt.Println(d)
			return
		}
		if ch >= 'a' && ch <= 'z' {
			if G[int(ch-'a')] != nil {
				for _, p := range G[int(ch-'a')] {
					q = append(q, pd{p, d + 1})
				}
				G[int(ch-'a')] = nil
			}
		}
		q = append(q, pd{pair{x - 1, y}, d + 1})
		q = append(q, pd{pair{x + 1, y}, d + 1})
		q = append(q, pd{pair{x, y - 1}, d + 1})
		q = append(q, pd{pair{x, y + 1}, d + 1})
	}
	fmt.Println(-1)
}

type pair struct{ x, y int }
type pd struct {
	p pair
	d int
}
