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

	p := make([]int, n)
	pos := make([]int, n)
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
		p[i]--
		pos[p[i]] = i
	}

	used := make([]bool, n-1)
	for i := 0; i < n; i++ {
		now := pos[i]
		for i < now {
			before := now - 1
			if used[before] {
				fmt.Fprintln(out, -1)
				return
			}
			used[before] = true
			pos[p[now]] = before
			pos[p[before]] = now
			p[before], p[now] = p[now], p[before]
			res = append(res, before)
			now--
		}
	}

	if len(res) != n-1 {
		fmt.Fprintln(out, -1)
	} else {
		for i := 0; i < n-1; i++ {
			fmt.Fprintln(out, res[i]+1)
		}
	}
}
