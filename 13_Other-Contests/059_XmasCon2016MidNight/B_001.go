package main

import "fmt"

func dfs(N, h, x, y int) {
	if h == N {
		fmt.Println(x, y)
		return
	}
	x += 1 << (N - h - 1)
	x++
	y++
	dfs(N, h+1, x, y)
	x -= 1 << (N - h - 1)
	x--
	y--
	fmt.Println(x, y)
	y += 1 << (N - h - 1)
	y++
	x++
	dfs(N, h+1, x, y)
	return
}

func main() {
	dfs(11, 0, 0, 0)
}
