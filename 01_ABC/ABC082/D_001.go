package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	var X, Y int
	fmt.Scan(&s, &X, &Y)

	dx, dy := make([]int, 0), make([]int, 0)
	d := 0
	s += "T"
	for i := range s {
		if s[i] == 'F' {
			d++
		} else {
			if len(dx) == len(dy) {
				dx = append(dx, d)
			} else {
				dy = append(dy, d)
			}
			d = 0
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(dx)))
	sort.Sort(sort.Reverse(sort.IntSlice(dy)))

	x, y := dx[len(dx)-1], 0
	dx = dx[:len(dx)-1]

	for i := range dx {
		if x < X {
			x += dx[i]
		} else {
			x -= dx[i]
		}
	}
	for i := range dy {
		if y < Y {
			y += dy[i]
		} else {
			y -= dy[i]
		}
	}

	if x == X && y == Y {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
