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

	var n, m int
	fmt.Fscan(in, &n, &m)

	s := make([]int, m)
	t := make([]int, m)
	a := make([]int, 1<<19)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &s[i], &t[i])
		a[s[i]]++
		a[t[i]+1]--
	}

	for i := 1; i < n+1; i++ {
		a[i] += a[i-1]
	}
	for i := 0; i < n+1; i++ {
		if a[i]-1 == 0 {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
	for i := 1; i < n+1; i++ {
		a[i] += a[i-1]
	}

	ans := make([]int, 0)
	for i := 0; i < m; i++ {
		if a[t[i]]-a[s[i]-1] == 0 {
			ans = append(ans, i+1)
		}
	}

	fmt.Fprintln(out, len(ans))
	for i := range ans {
		fmt.Fprintln(out, ans[i])
	}
}
