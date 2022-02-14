package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	mi := 1
	a := make([]int, n+1)
	s := [2]int{}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		s[i%2] += max(a[i], 0)
		if a[i] > a[mi] {
			mi = i
		}
	}

	z := 0
	cnt := 0
	tmp := 0
	if s[1] > s[0] {
		tmp = 1
	}
	x := make([]int, 505)
	y := make([]int, 505)
	for i := tmp; i <= n; i += 2 {
		if a[i] > 0 {
			z++
			x[z] = i
			if z > 1 {
				y[z-1] = (x[z] - x[z-1]) / 2
				cnt += y[z-1]
			}
		}
	}

	var ans int
	if z != 0 {
		ans = max(s[0], s[1])
		cnt += x[1] - 1 + n - x[z]
	} else {
		ans = a[mi]
		cnt = n - 1
		z++
		x[z] = mi
	}
	fmt.Fprintln(out, ans)
	fmt.Fprintln(out, cnt)

	for i := n; i > x[z]; i-- {
		fmt.Fprintln(out, i)
	}
	for i := 1; i < x[1]; i++ {
		fmt.Fprintln(out, 1)
	}
	for i := 1; i < z; i++ {
		for j := y[i]; j > 0; j-- {
			fmt.Fprintln(out, j+1)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
