package main

import "fmt"

func main() {
	var x, y int
	var w string
	fmt.Scan(&x, &y, &w)
	x--
	y--

	c := make([]string, 9)
	for i := range c {
		fmt.Scan(&c[i])
	}

	n := len(w)
	dx, dy := 0, 0
	if w[0] == 'R' {
		dx = 1
	}
	if w[0] == 'L' {
		dx = -1
	}
	if w[n-1] == 'U' {
		dy = -1
	}
	if w[n-1] == 'D' {
		dy = 1
	}

	for i := 0; i < 4; i++ {
		fmt.Print(string(c[y][x]))
		next_x := x + dx
		next_y := y + dy
		if next_x < 0 || next_x >= 9 {
			dx *= -1
		}
		if next_y < 0 || next_y >= 9 {
			dy *= -1
		}
		x += dx
		y += dy
	}
	fmt.Println()
}
