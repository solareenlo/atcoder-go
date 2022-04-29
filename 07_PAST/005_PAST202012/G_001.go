package main

import (
	"fmt"
	"strings"
)

type pair struct{ x, y int }

var (
	h, w int
	B    = make([][]string, 4)
	ans  = make([]pair, 0)
	tmp  = make([]pair, 0)
	dx   = []int{0, -1, 0, 1}
	dy   = []int{1, 0, -1, 0}
)

func dfs(i, j int) bool {
	B[i][j] = "."
	tmp = append(tmp, pair{i, j})
	ok := true
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if B[i][j] == "#" {
				ok = false
			}
		}
	}
	if ok {
		ans = tmp
	}
	for k := 0; k < 4; k++ {
		if ok {
			break
		}
		x := i + dx[k]
		y := j + dy[k]
		if 0 <= x && x < h && 0 <= y && y < w && B[x][y] == "#" {
			if dfs(x, y) {
				ok = true
			}
		}
	}
	B[i][j] = "#"
	tmp = tmp[:len(tmp)-1]
	return ok
}

func main() {
	fmt.Scan(&h, &w)
	for i := 0; i < h; i++ {
		var b string
		fmt.Scan(&b)
		B[i] = strings.Split(b, "")
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if B[i][j] == "#" && dfs(i, j) {
				fmt.Println(len(ans))
				for i := range ans {
					fmt.Println(ans[i].x+1, ans[i].y+1)
				}
				return
			}
		}
	}
}
