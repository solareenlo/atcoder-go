package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct {
	x, y int
}

func calc(p1, p2, p3 point) int {
	v1 := point{p2.x - p1.x, p2.y - p1.y}
	v2 := point{p3.x - p1.x, p3.y - p1.y}
	return v1.x*v2.y - v1.y*v2.x
}

func cross(a1, a2, b1, b2 point) bool {
	return calc(a1, a2, b1)*calc(a1, a2, b2) < 0 && calc(b1, b2, a1)*calc(b1, b2, a2) < 0
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, Q, tmp0, tmp1 int
	fmt.Fscan(in, &N, &M, &Q, &tmp0, &tmp1)
	var p [255]point
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &p[i].x, &p[i].y)
	}
	V := N + Q
	var s [55][6]point
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &s[i][0].x, &s[i][0].y, &s[i][1].x, &s[i][1].y)
		s[i][2] = point{s[i][0].x, s[i][1].y}
		s[i][3] = point{s[i][1].x, s[i][0].y}
		for j := 0; j < 4; j++ {
			p[V].x = s[i][j].x
			p[V].y = s[i][j].y
			V++
		}
	}
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &p[i+N].x, &p[i+N].y)
	}
	var d [255][255]float64
	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			d[i][j] = math.Hypot(float64(p[i].x-p[j].x), float64(p[i].y-p[j].y))
			for k := 0; k < M; k++ {
				for l := 0; l < 4; l += 2 {
					if cross(p[i], p[j], s[k][l], s[k][l+1]) {
						d[i][j] = 1.0e+9
					}
				}
			}
		}
	}
	sum := 0.0
	var dist [255]float64
	var vis [255]bool
	for i := 0; i < Q; i++ {
		for j := 0; j < V; j++ {
			dist[j] = 1.0e+9
			vis[j] = false
		}
		dist[i+N] = 0.0
		for j := 0; j < V; j++ {
			cur := 1.0e+9
			pos := -1
			for k := 0; k < V; k++ {
				if !vis[k] && dist[k] < cur {
					cur = dist[k]
					pos = k
				}
			}
			vis[pos] = true
			for k := 0; k < V; k++ {
				if dist[k] > dist[pos]+d[pos][k] {
					dist[k] = dist[pos] + d[pos][k]
				}
			}
		}
		ret := 1.0e+9
		for j := 0; j < N; j++ {
			if dist[j] < ret {
				ret = dist[j]
			}
		}
		sum += ret * 2
	}
	fmt.Println(sum)
}
