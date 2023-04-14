package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s, t string
	fmt.Fscan(in, &s, &t)
	sz := len(s)
	var lc [5005][5005]int
	for i := 1; i <= sz; i++ {
		for j := 1; j <= sz; j++ {
			if s[i-1] == t[j-1] {
				lc[i][j] = lc[i-1][j-1] + 1
			} else {
				lc[i][j] = max(lc[i-1][j], lc[i][j-1])
			}
		}
	}
	fmt.Println(lc[sz][sz] + 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
