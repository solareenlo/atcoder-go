package main

import "fmt"

var (
	dx   = [4]int{0, 1, 0, -1}
	dy   = [4]int{1, 0, -1, 0}
	h, w int
	s    = [404]string{}
	seen = [404][404]bool{}
	c    = [2]int{}
)

func dfs(i, j int, ng byte) {
	if i < 0 || h <= i || j < 0 || w <= j {
		return
	}
	if s[i][j] == ng {
		return
	}
	if seen[i][j] {
		return
	}
	seen[i][j] = true
	if s[i][j] == '#' {
		c[0]++
	} else {
		c[1]++
	}
	for k := 0; k < 4; k++ {
		dfs(i+dy[k], j+dx[k], s[i][j])
	}
}

func main() {
	fmt.Scan(&h, &w)

	for i := 0; i < h; i++ {
		fmt.Scan(&s[i])
	}

	res := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			c[0], c[1] = 0, 0
			dfs(i, j, '0')
			res += c[0] * c[1]
		}
	}
	fmt.Println(res)
}
