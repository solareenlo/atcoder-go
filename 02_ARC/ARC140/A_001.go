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

	var n, x int
	var s string
	fmt.Fscan(in, &n, &x, &s)

	for i := 1; i <= n; i++ {
		if n%i != 0 {
			continue
		}
		cnt := 0
		for j := 0; j < i; j++ {
			mp := make(map[byte]int)
			maxx := 0
			for k := j; k < n; k += i {
				mp[s[k]]++
				maxx = max(maxx, mp[s[k]])
			}
			cnt += n/i - maxx
		}
		if cnt <= x {
			fmt.Fprintln(out, i)
			break
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
