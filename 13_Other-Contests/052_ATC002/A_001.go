package main

import (
	"fmt"
)

type pair struct {
	x, y int
}

func main() {
	var H, W, sy, sx, gy, gx int
	fmt.Scan(&H, &W, &sy, &sx, &gy, &gx)
	sy--
	sx--
	gy--
	gx--
	s := make([]string, H)
	for i := range s {
		fmt.Scan(&s[i])
	}
	dir := []int{1, 0, -1, 0, 1}
	dis := make([][]int, H)
	for i := range dis {
		dis[i] = make([]int, W)
		for j := range dis[i] {
			dis[i][j] = 1e9
		}
	}
	dis[sy][sx] = 0
	Q := []pair{{sy, sx}}
	for len(Q) > 0 {
		cy, cx := Q[0].x, Q[0].y
		Q = Q[1:]
		for i := 0; i < 4; i++ {
			ny, nx := cy+dir[i], cx+dir[i+1]
			if ny < 0 || nx < 0 || ny >= H || nx >= W {
				continue
			}
			if s[ny][nx] == '#' {
				continue
			}
			if dis[ny][nx] > dis[cy][cx]+1 {
				dis[ny][nx] = dis[cy][cx] + 1
				Q = append(Q, pair{ny, nx})
			}
		}
	}
	fmt.Println(dis[gy][gx])
}
