package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	type pair struct{ x, y int }
	p := make([]pair, 0)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		p = append(p, pair{x, y})
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].x < p[j].x
	})

	s := map[int]bool{}
	s[n] = true
	l := 0
	for l < m {
		a := make([]int, 0)
		b := make([]int, 0)
		r := l
		for r < m {
			if p[r].x == p[l].x {
				r++
			} else {
				break
			}
		}
		for i := l; i < r; i++ {
			y := p[i].y
			_, ok0 := s[y-1]
			_, ok1 := s[y+1]
			_, ok2 := s[y]
			if (ok0 || ok1) && !ok2 {
				a = append(a, y)
			}
			if (!ok0 && !ok1) && ok2 {
				b = append(b, y)
			}
		}
		for i := 0; i < len(a); i++ {
			s[a[i]] = true
		}
		for i := 0; i < len(b); i++ {
			delete(s, b[i])
		}
		l = r
	}

	fmt.Println(len(s))
}
