package main

import (
	"fmt"
	"strings"
)

var (
	dx    = [4]int{1, 0, -1, 0}
	dy    = [4]int{0, 1, 0, -1}
	field = make([]string, 10)
	seen  = [12][12]bool{}
)

func dfs(h, w int) {
	seen[h][w] = true
	for i := 0; i < 4; i++ {
		nh := h + dx[i]
		nw := w + dy[i]
		if nh < 0 || nh >= 10 || nw < 0 || nw >= 10 {
			continue
		}
		if field[nh][nw] == 'x' {
			continue
		}
		if seen[nh][nw] {
			continue
		}
		dfs(nh, nw)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Scan(&field[i])
	}

	sum := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if field[i][j] == 'o' {
				sum++
			}
		}
	}

	ok := false
	change := false
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if field[i][j] == 'x' {
				tmp := strings.Split(field[i], "")
				tmp[j] = "o"
				field[i] = strings.Join(tmp, "")
				change = true
			}
			for k := range seen {
				for l := range seen[k] {
					seen[k][l] = false
				}
			}
			dfs(i, j)
			cnt := 0
			for h := 0; h < 10; h++ {
				for w := 0; w < 10; w++ {
					if field[h][w] == 'o' && seen[h][w] {
						cnt++
					}
				}
			}
			if cnt == sum+1 {
				ok = true
			}
			if change {
				tmp := strings.Split(field[i], "")
				tmp[j] = "x"
				field[i] = strings.Join(tmp, "")
				change = false
			}
		}
	}

	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
