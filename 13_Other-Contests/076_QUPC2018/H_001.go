package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var n int
var a []int
var ch [200005][26]bool

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	fmt.Println(solve())
}

func solve() string {
	id := 0
	ans := make([]string, 0)
	for i := 0; i < n; i++ {
		ans = append(ans, "0")
	}
	for i := 0; i < n; i++ {
		now := 0
		r := a[i]/2 + 1
		if ans[i] == "0" {
			for ch[i][now] {
				now++
			}
			ans[i] = string('a' + now)
		}
		for j := max(1, id-i+1); j < r; j++ {
			ans[i+j] = ans[i-j]
			id = i + j
		}
		id = max(id, i)
		if i-r >= 0 && i+r < n {
			ch[i+r][ans[i-r][0]-'a'] = true
		}
	}
	return strings.Join(ans, "")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
