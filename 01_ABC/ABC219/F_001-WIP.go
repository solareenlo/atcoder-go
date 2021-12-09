package main

import "fmt"

func main() {
	var s string
	var k int
	fmt.Scan(&s, &k)

	type pair struct{ x, y int }
	st := map[pair]bool{}

	x, y := 0, 0
	for _, c := range s {
		st[pair{x, y}] = true
		switch c {
		case 'L':
			x--
		case 'R':
			x++
		case 'U':
			y--
		case 'D':
			y++
		}
	}
	st[pair{x, y}] = true

	if x == 0 && y == 0 {
		fmt.Println(len(st))
		return
	}

	ps := make([]pair, 0)
	for p := range st {
		ps = append(ps, p)
	}

	if x == 0 {
		x, y = y, x
		for _, p := range ps {
			p.x, p.y = p.y, p.x
		}
	}
	if x < 0 {
		x = -x
		for _, p := range ps {
			p.x *= -1
		}
	}
	if y < 0 {
		y = -y
		for _, p := range ps {
			p.y *= -1
		}
	}
}
