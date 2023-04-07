package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pair struct {
	x, y int
}

type P struct {
	x pair
	y int
}

var h, w, x int
var m [1010][]string
var que []P
var ans int = -1
var dx [4]int = [4]int{-1, 0, 1, 0}
var dy [4]int = [4]int{0, 1, 0, -1}

func main() {
	in := bufio.NewReader(os.Stdin)

	var s pair
	fmt.Fscan(in, &h, &w, &x)
	for i := 0; i < h; i++ {
		var t string
		fmt.Fscan(in, &t)
		m[i] = strings.Split(t, "")
		for j := 0; j < w; j++ {
			if m[i][j] == "@" {
				que = append(que, P{pair{i, j}, 0})
			}
			if m[i][j] == "S" {
				s = pair{i, j}
			}
		}
	}
	bfs(x)
	que = append(que, P{s, 0})
	bfs(1000100)
	fmt.Println(ans)
}

func bfs(z int) {
	for len(que) > 0 {
		p := que[0]
		x := p.x.x
		y := p.x.y
		d := p.y
		que = que[1:]
		if m[x][y] == "#" {
			continue
		}
		if z == 1000100 && m[x][y] == "G" {
			ans = d
			break
		}
		m[x][y] = "#"
		if d == z {
			continue
		}
		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			if m[nx][ny] == "#" || m[nx][ny] == "@" {
				continue
			}
			que = append(que, P{pair{nx, ny}, d + 1})
		}
	}
}
