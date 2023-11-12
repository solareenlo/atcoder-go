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

	var H, W int
	fmt.Fscan(in, &H, &W)
	s := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Fscan(in, &s[i])
	}

	var st, at, gt int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			t := i*W + j
			if s[i][j] == 's' {
				st = t
			}
			if s[i][j] == 'a' {
				at = t
			}
			if s[i][j] == 'g' {
				gt = t
			}
		}
	}
	dy := [4]int{1, -1, 0, 0}
	dx := [4]int{0, 0, 1, -1}
	q := make([]P, 0)
	v := make([][]int, H*W)
	for i := range v {
		v[i] = make([]int, H*W)
		for j := range v[i] {
			v[i][j] = -1
		}
	}
	q = append(q, P{st, at})
	v[st][at] = 0
	for len(q) > 0 {
		u, a := q[0].x, q[0].y
		q = q[1:]
		d := v[u][a]
		y := u / W
		x := u % W
		if a == gt {
			fmt.Println(v[u][a])
			return
		}
		for i := 0; i < 4; i++ {
			ny := y + dy[i]
			nx := x + dx[i]
			if ny < 0 || ny >= H || nx < 0 || nx >= W {
				continue
			}
			if s[ny][nx] == '#' {
				continue
			}
			t := ny*W + nx
			aa := a
			if t == a {
				ry := ny + dy[i]
				rx := nx + dx[i]
				if ry < 0 || ry >= H || rx < 0 || rx >= W {
					continue
				}
				if s[ry][rx] == '#' {
					continue
				}
				tt := ry*W + rx
				aa = tt
			}
			if v[t][aa] >= 0 {
				continue
			}
			v[t][aa] = d + 1
			q = append(q, P{t, aa})
		}
	}
	fmt.Println(-1)
}
