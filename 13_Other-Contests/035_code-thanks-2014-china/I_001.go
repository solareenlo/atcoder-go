package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type P struct {
		x, y int
	}
	vy := []int{1, 0, -1, 0}
	vx := []int{0, 1, 0, -1}

	var N int
	fmt.Fscan(in, &N)
	S := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &S[i])
	}

	que := make([]P, 0)
	que = append(que, P{N - 1, N - 1})
	var v [1000][1000]bool
	v[N-1][N-1] = true
	var sz [1000][1000]int
	sz[N-1][N-1] = 2
	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		for i := 0; i < 4; i++ {
			ny := p.x + vy[i]
			nx := p.y + vx[i]
			if ny < 0 || nx < 0 || ny >= N || nx >= N {
				continue
			}
			if v[ny][nx] {
				continue
			}
			sz[ny][nx]++
			if S[p.x][p.y] == '#' {
				sz[ny][nx] += 2
			}
			if sz[ny][nx] >= 2 {
				v[ny][nx] = true
				que = append(que, P{ny, nx})
			}
		}
	}
	if v[0][0] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
