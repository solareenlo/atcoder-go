package main

import (
	"fmt"
	"sort"
)

func main() {
	type point struct {
		x, y int
		lock bool
	}

	var n int
	fmt.Scan(&n)

	red, blue := make([]point, n), make([]point, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&red[i].y, &red[i].x)
	}
	sort.Slice(red, func(i, j int) bool {
		return red[i].y > red[j].y
	})
	for i := 0; i < n; i++ {
		fmt.Scan(&blue[i].y, &blue[i].x)
	}
	sort.Slice(blue, func(i, j int) bool {
		return blue[i].x < blue[j].x
	})

	cnt := 0
	for i := range blue {
		for j := range red {
			if red[j].x < blue[i].x && red[j].y < blue[i].y && !red[j].lock {
				red[j].lock = true
				cnt++
				break
			}
		}
	}
	fmt.Println(cnt)
}
