package main

import "fmt"

var dx [4]int = [4]int{1, 0, -1, 0}
var dy [4]int = [4]int{0, 1, 0, -1}
var d [5][16][16][16][16]bool

type p struct{ d, a, b, x, y int }

func main() {
	var a, b, h, w int
	fmt.Scan(&a, &b, &h, &w)

	c := make([]string, h)
	for i := range c {
		fmt.Scan(&c[i])
	}

	q := make([]p, 0)
	q = append(q, p{0, a, b, 1, 1})
	d[0][1][1][a][b] = true
	for len(q) != 0 {
		t := q[0]
		q = q[1:]
		if t.x == h-2 && t.y == w-2 && t.a == 0 && t.b == 0 {
			fmt.Println("Yes")
			return
		}
		for i := 3; i < 6; i++ {
			s := t
			s.d = (t.d + i) % 4
			s.x = t.x + dx[s.d]
			s.y = t.y + dy[s.d]
			if i > 4 {
				s.a -= 1
			}
			if i < 4 {
				s.b -= 1
			}
			if c[s.x][s.y] == '#' || s.a < 0 || s.b < 0 || d[s.d][s.a][s.b][s.x][s.y] {
				continue
			}
			q = append(q, s)
			d[s.d][s.a][s.b][s.x][s.y] = true
		}
	}
	fmt.Println("No")
}
