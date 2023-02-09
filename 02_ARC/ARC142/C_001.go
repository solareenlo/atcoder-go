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

	ans := int(1e18)

	var n int
	fmt.Fscan(in, &n)

	flag := 1
	var x, y int
	for i := 3; i <= n; i++ {
		fmt.Fprintf(out, "? 1 %d\n", i)
		out.Flush()
		var a int
		fmt.Fscan(in, &a)
		fmt.Fprintf(out, "? 2 %d\n", i)
		out.Flush()
		var b int
		fmt.Fscan(in, &b)
		ans = min(ans, a+b)
		if abs(b-a) == 1 {
			flag &= 1
		} else {
			flag &= 0
		}
		if a+b == 3 && x != 0 {
			y = i
		} else if a+b == 3 {
			x = i
		}
	}
	if x != 0 && y != 0 {
		fmt.Fprintf(out, "? %d %d\n", x, y)
		out.Flush()
		var a int
		fmt.Fscan(in, &a)
		if a == 1 {
			fmt.Fprintf(out, "! 3\n")
			out.Flush()
			os.Exit(0)
		}
	}
	if flag == 0 {
		fmt.Fprintf(out, "! %d\n", ans)
		out.Flush()
	} else {
		fmt.Fprintf(out, "! 1\n")
		out.Flush()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
