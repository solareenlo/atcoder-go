package main

import (
	"fmt"
	"sort"
)

func main() {
	type point struct{ x, y int }

	var n int
	fmt.Scan(&n)

	red, blue := make([]point, n), make([]point, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&red[i].x, &red[i].y)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&blue[i].x, &blue[i].y)
	}
	sort.Slice(blue, func(i, j int) bool {
		return blue[i].x < blue[j].x
	})

	cnt := 0
	for i := range blue {
		hit := point{-1, -1}
		index := -1
		for j := range red {
			if red[j].x < blue[i].x && red[j].y < blue[i].y && hit.y < red[j].y {
				hit = red[j]
				index = j
			}
		}
		if index != -1 {
			cnt++
			red = append(red[:index], red[index+1:]...)
		}
	}
	fmt.Println(cnt)
}
