package main

import "fmt"

func main() {
	var dx = [4]int{0, 1, 0, -1}
	var dy = [4]int{1, 0, -1, 0}

	var n int
	fmt.Scan(&n)

	if n >= 10000 {
		n = 10000 + (n-10000)%104
	}
	x := 1000
	y := 1000
	var v [2000][2000]int
	st := 0
	for i := 0; i < n; i++ {
		if v[y][x] != 0 {
			st = (st + 1) % 4
		} else {
			st = (st + 3) % 4
		}
		v[y][x] = 1 - v[y][x]
		x += dx[st]
		y += dy[st]
	}
	fmt.Println(v[y][x])
}
