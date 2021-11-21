package main

import "fmt"

var (
	H, W, RES int
	MAP       = [16][16]bool{}
)

func dfs(y, x, a, b int) {
	if x == W {
		y++
		x = 0
	}
	if y == H {
		RES++
		return
	}
	if MAP[y][x] {
		dfs(y, x+1, a, b)
		return
	}
	MAP[y][x] = true
	if b != 0 {
		dfs(y, x+1, a, b-1)
	}
	if a != 0 && MAP[y][x+1] == false {
		MAP[y][x+1] = true
		dfs(y, x+1, a-1, b)
		MAP[y][x+1] = false
	}
	if a != 0 && MAP[y+1][x] == false {
		MAP[y+1][x] = true
		dfs(y, x+1, a-1, b)
		MAP[y+1][x] = false
	}
	MAP[y][x] = false
}

func main() {
	var a, b int
	fmt.Scan(&H, &W, &a, &b)

	dfs(0, 0, a, b)
	fmt.Println(RES)
}
