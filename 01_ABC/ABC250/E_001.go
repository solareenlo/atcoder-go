package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)

func hread(psum []int, n int) {
	s := make(map[int]bool)
	for i := 1; i <= n; i++ {
		psum[i] = psum[i-1]
		var x int
		fmt.Fscan(in, &x)
		if _, ok := s[x]; !ok {
			psum[i] += x * (x + 93) * (x + 117)
		}
		s[x] = true
	}
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	suma := make([]int, n+1)
	sumb := make([]int, n+1)
	hread(suma, n)
	hread(sumb, n)

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if suma[x] == sumb[y] {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
