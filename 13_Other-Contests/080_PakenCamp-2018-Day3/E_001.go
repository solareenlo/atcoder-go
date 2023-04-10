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

	var q int
	fmt.Fscan(in, &q)
	for j := 0; j < q; j++ {
		a := make([]int, 2)
		var x int
		fmt.Fscan(in, &a[0], &a[1], &x)
		for i := 1; ; i++ {
			if a[i-1]+a[i] > x {
				break
			}
			a = append(a, a[i-1]+a[i])
		}
		n := len(a)
		ans := 0
		for k := 0; k < 2; k++ {
			var y int
			if k == 0 {
				y = x
			} else {
				y = x - a[0]
			}
			for i := n - 1; i >= k+1; i-- {
				if y >= a[i] {
					y -= a[i]
				}
			}
			if y == 0 {
				ans++
			}
		}
		fmt.Fprintln(out, ans)
	}
}
