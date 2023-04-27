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

	var f [100]int
	f[0] = 1
	f[1] = 1
	for i := 2; f[i-1]+f[i-2] < int(1e18); i++ {
		f[i] = f[i-1] + f[i-2]
	}

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		var n int
		fmt.Fscan(in, &n)
		ans := 0
		for i := 0; i < n; i++ {
			var x int
			fmt.Fscan(in, &x)
			p := 1
			for j := 86; j > 0; j-- {
				if f[j] <= x {
					x -= f[j]
					p = j
				}
			}
			if p == 1 {
				ans ^= 1
			} else {
				ans ^= (p & 1) * 2
			}
		}
		if ans != 0 {
			fmt.Fprintln(out, "Black")
		} else {
			fmt.Fprintln(out, "White")
		}
	}
}
