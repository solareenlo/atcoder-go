package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var r, c, sy, sx, gy, gx int
	fmt.Fscan(in, &r, &c, &sy, &sx, &gy, &gx)
	sy--
	sx--
	gy--
	gx--
	C := make([]string, r)
	for i := range C {
		fmt.Fscan(in, &C[i])
	}
	dist := make([][]int, r)
	for i := range dist {
		dist[i] = make([]int, c)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[sy][sx] = 0
	Qi := make([]int, 0)
	Qj := make([]int, 0)
	Qi = append(Qi, sy)
	Qj = append(Qj, sx)
	di := []int{1, 0, -1, 0}
	dj := []int{0, 1, 0, -1}
	for len(Qi) != 0 {
		i := Qi[0]
		j := Qj[0]
		Qi = Qi[1:]
		Qj = Qj[1:]
		for k := 0; k < 4; k++ {
			ni := i + di[k]
			nj := j + dj[k]
			if 0 <= ni && ni < r && 0 <= nj && nj < c && C[ni][nj] == '.' && dist[ni][nj] == -1 {
				dist[ni][nj] = dist[i][j] + 1
				Qi = append(Qi, ni)
				Qj = append(Qj, nj)
			}
		}
	}
	fmt.Println(dist[gy][gx])
}
