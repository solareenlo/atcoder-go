package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	var a, p [105]int
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &a[i])
	}

	mx := 0
	y := make([][]int, 105)
	for i := range y {
		y[i] = make([]int, 105)
	}
	for i := 1; i <= n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < m; j++ {
			if s[j] == 'o' {
				p[i] += a[j+1]
			} else {
				y[i][0]++
				y[i][y[i][0]] = a[j+1]
			}
		}
		p[i] += i
		mx = max(mx, p[i])
		sort.Ints(y[i][1 : y[i][0]+1])
	}
	for i := 1; i <= n; i++ {
		cnt := 0
		for p[i] < mx {
			p[i] += y[i][y[i][0]]
			y[i][0]--
			cnt++
		}
		fmt.Fprintln(out, cnt)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
