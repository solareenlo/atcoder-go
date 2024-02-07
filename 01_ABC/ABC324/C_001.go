package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)
	d := len(s)
	var b [500010]bool
	k := 0
	for i := 1; i <= n; i++ {
		var a string
		fmt.Fscan(in, &a)
		t := len(a)
		cnt := 0
		if abs(t-d) > 1 {
			continue
		}
		for x, y := 0, 0; x < d && y < t; x, y = x+1, y+1 {
			if s[x] != a[y] {
				cnt++
				if d > t {
					y--
				} else if d < t {
					x--
				}
			}
		}
		if cnt <= 1 {
			b[i] = true
			k++
		}
	}
	fmt.Println(k)
	for i := 1; i <= n; i++ {
		if b[i] {
			fmt.Printf("%d ", i)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
