package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)

	wx, wy, wc := 0, 0, 0
	ans := 0
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if a == x {
			wx = i
		}
		if a == y {
			wy = i
		}
		if a < y || a > x {
			wc = i
		}
		if wc < wx && wc < wy {
			ans += min(wx, wy) - wc
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
