package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	ans := 0
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		if i > 1 {
			ans += (a[i] - a[i-1] + m) % m
		}
	}

	d := make([]int, 200200)
	d2 := make([]int, 200200)
	for i := 1; i < n; i++ {
		x := a[i]
		y := a[i+1]
		d2[x+2]++
		d2[y+1]--
		if x < y {
			d[y+1] -= y - x - 1
		} else {
			d[y+1] -= m + y - x - 1
			d2[1] += m - x
			d2[2] += x - m + 1
		}
	}

	t := 0
	for i := 1; i <= m; i++ {
		d2[i] += d2[i-1]
		d[i] += d[i-1] + d2[i]
		t = max(t, d[i])
	}
	fmt.Println(ans - t)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
