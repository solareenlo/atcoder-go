package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var dy [4]int = [4]int{1, -1, 0, 0}
var dz [4]int = [4]int{0, 0, 1, -1}

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)
	var dp [12][12][12][110]int
	var d2 [12][12][12][110]int
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			for k := 0; k < 12; k++ {
				for l := 0; l < 110; l++ {
					dp[i][j][k][l] = INF
					d2[i][j][k][l] = INF
				}
			}
		}
	}
	var t, w, h, lev int
	fmt.Fscan(in, &t, &w, &h, &lev)
	var ma [12][12]string
	ma[0][0], _ = in.ReadString('\n')
	for i := 0; i < t; i++ {
		for j := 0; j < h; j++ {
			ma[i][j], _ = in.ReadString('\n')
		}
	}
	var fe [12][12][12]int
	var sx, sy, sz, gx, gy, gz int
	for i := 0; i < t; i++ {
		for j := 0; j < h; j++ {
			for k := 0; k < w; k++ {
				if ma[i][j][k*2] >= '0' && ma[i][j][k*2] <= '9' {
					fe[i][j][k] = int(ma[i][j][k*2]-'0')*10 + int(ma[i][j][k*2+1]-'0')
				}
				if ma[i][j][k*2] == ' ' || ma[i][j][k*2] == '#' {
					fe[i][j][k] = -3
				}
				if ma[i][j][k*2] == '_' {
					fe[i][j][k] = -1
				}
				if ma[i][j][k*2] == '-' {
					fe[i][j][k] = -2
				}
				if ma[i][j][k*2] == 'K' {
					sx = i
					sy = j
					sz = k
				}
				if ma[i][j][k*2] == '$' {
					gx = i
					gy = j
					gz = k
				}
				if ma[i][j][k*2] == 'H' {
					fe[i][j][k] = 100
				}
			}
		}
	}

	q := &Heap{}

	var aedge func(int, int, int, int, int)
	aedge = func(x, y, z, le, ti int) {
		le = min(le, 100)
		if x < 0 || x >= t || y < 0 || y >= h || z < 0 || z >= w {
			return
		}
		if dp[x][y][z][le] <= ti {
			return
		}
		dp[x][y][z][le] = ti
		heap.Push(q, IPP{-ti, PP{P{x, y}, P{z, le}}})
		return
	}

	aedge(sx, sy, sz, lev, 0)
	for q.Len() > 0 {
		p := heap.Pop(q).(IPP)
		ti := -p.x
		x := p.y.x.x
		y := p.y.x.y
		z := p.y.y.x
		le := p.y.y.y
		if d2[x][y][z][le] < INF {
			continue
		}
		d2[x][y][z][le] = ti
		if x == gx && y == gy && z == gz {
			fmt.Println(ti)
			return
		}
		if fe[x][y][z] == -1 {
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if fe[x+1][i][j] == -2 {
						aedge(x+1, i, j, le, ti+1)
					}
				}
			}
		}
		if fe[x][y][z] == -2 && x > sx {
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if fe[x-1][i][j] == -1 {
						aedge(x-1, i, j, le, ti+1)
					}
				}
			}
		}
		for i := 0; i < 4; i++ {
			nz := z + dz[i]
			ny := y + dy[i]
			if ny < 0 || nz < 0 || ny >= h || nz >= w {
				continue
			}
			if fe[x][ny][nz] < -2 || fe[x][ny][nz] > le {
				continue
			}
			if fe[x][ny][nz] == le {
				aedge(x, ny, nz, le+1, ti+1+fe[x][ny][nz])
			} else if fe[x][ny][nz] <= 0 {
				aedge(x, ny, nz, le, ti+1)
			} else {
				aedge(x, ny, nz, le, ti+1+fe[x][ny][nz])
			}
		}
	}
	fmt.Println("Impossible")
}

type P struct {
	x, y int
}

type PP struct {
	x, y P
}

type IPP struct {
	x int
	y PP
}

type Heap []IPP

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(IPP)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
