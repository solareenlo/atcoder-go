package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const N = 40000

var m [N][]string
var h, w int
var bx, by, a int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &h, &w)
	for i := 0; i < h; i++ {
		var t string
		fmt.Fscan(in, &t)
		m[i] = strings.Split(t, "")
		for j := 0; j < w; j++ {
			if m[i][j] == "S" {
				bx = i
				by = j
			}
		}
	}
	dfs(bx, by, 0)
	if a != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func dfs(x, y, s int) {
	if x == bx && y == by && s > 2 {
		a = 1
		return
	}
	if x < 0 || x >= h || y < 0 || y >= w || m[x][y] == "#" {
		return
	}
	m[x][y] = "#"
	dfs(x+1, y, s+1)
	dfs(x, y+1, s+1)
	dfs(x-1, y, s+1)
	dfs(x, y-1, s+1)
}
