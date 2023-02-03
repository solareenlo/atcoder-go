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

	const N = 200005

	var n int
	fmt.Fscan(in, &n)

	var a [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	tmp := a[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	t := n
	var s, c int
	for i := 0; i < n; i++ {
		z := 0
		for s <= i && t > 0 {
			c = 0
			for a[t] == a[t-1] && t > 1 {
				c++
				t--
			}
			if t > 0 {
				c++
				t--
			}
			if s == i {
				z += c
			}
			if t > 0 {
				s++
			}
		}
		fmt.Fprintln(out, z)
	}
}
