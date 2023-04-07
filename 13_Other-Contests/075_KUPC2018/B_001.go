package main

import (
	"fmt"
	"os"
	"strings"
)

var h, w int
var B [10]string
var ans []string

func main() {
	fmt.Scan(&h, &w)
	for i := 0; i < h; i++ {
		fmt.Scan(&B[i])
	}

	var x int
	for j := 0; j < w; j++ {
		if B[h-1][j] == 's' {
			x = j
		}
	}

	ans = make([]string, 11)
	dfs(0, x)
	fmt.Println("impossible")
}

func dfs(t, x int) {
	if B[h-1-t][x] == 'x' {
		return
	}
	if t == h-1 {
		fmt.Println(strings.Join(ans, ""))
		os.Exit(0)
	}
	ans[t] = "S"
	dfs(t+1, x)
	if x > 0 {
		ans[t] = "L"
		dfs(t+1, x-1)
	}
	if x < w-1 {
		ans[t] = "R"
		dfs(t+1, x+1)
	}
}
