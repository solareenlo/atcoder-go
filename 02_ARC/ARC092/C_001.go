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

	red := make([]P, n)
	blue := make([]P, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &red[i].x, &red[i].y)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &blue[i].x, &blue[i].y)
	}
	sort.Slice(blue, func(i, j int) bool {
		return blue[i].x < blue[j].x
	})

	res := 0
	for i := 0; i < n; i++ {
		hit := P{-1, -1}
		var index int
		for j := 0; j < len(red); j++ {
			if red[j].x < blue[i].x && red[j].y < blue[i].y && hit.y < red[j].y {
				hit = red[j]
				index = j
			}
		}
		if hit.x != -1 && hit.y != -1 {
			res++
			red = erase(red, index)
		}
	}
	fmt.Println(res)
}

type P struct{ x, y int }

func erase(a []P, pos int) []P {
	return append(a[:pos], a[pos+1:]...)
}
