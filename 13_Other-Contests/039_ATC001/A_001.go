package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)

	a := make([][]string, h+2)
	for i := range a {
		a[i] = strings.Split(strings.Repeat("#", w+2), "")
	}
	p := 0
	q := 0
	for i := 1; i <= h; i++ {
		var s string
		fmt.Fscan(in, &s)
		s = "#" + s + strings.Repeat("#", w-len(s)+1)
		a[i] = strings.Split(s, "")
		for j := 1; j <= w; j++ {
			if a[i][j] == "s" {
				p = j
				q = i
			}
		}
	}

	ok := false
	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 1 || y < 1 || x > w || y > h || a[y][x] == "#" {
			return
		}
		if a[y][x] == "g" {
			ok = true
		}
		a[y][x] = "#"
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	dfs(p, q)

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
