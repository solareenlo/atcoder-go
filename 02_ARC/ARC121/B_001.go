package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([][]int, 3)
	rgb := "RGB"
	for i := 0; i < 2*n; i++ {
		var x int
		var c string
		fmt.Fscan(in, &x, &c)
		idx := strings.IndexByte(rgb, c[0])
		a[idx] = append(a[idx], x)
	}

	res := make([]int, 0)
	for i := 0; i < 3; i++ {
		sort.Ints(a[i])
		if len(a[i])%2 != 0 {
			res = append(res, i)
		}
	}
	if len(res) == 0 {
		fmt.Println(0)
		return
	}

	u := res[0]
	v := res[1]
	cal := func(u, v int) int {
		ans := 1 << 60
		if len(a[v]) == 0 {
			return ans
		}
		j := 0
		for i := 0; i < len(a[u]); i++ {
			for j+1 < len(a[v]) && a[v][j+1] <= a[u][i] {
				j++
			}
			ans = min(ans, abs(a[u][i]-a[v][j]))
			if j+1 < len(a[v]) {
				ans = min(ans, abs(a[u][i]-a[v][j+1]))
			}
		}
		return ans
	}

	w := 3 - u - v
	fmt.Println(min(cal(u, v), cal(u, w)+cal(v, w)))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
