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

	var t int
	fmt.Fscan(in, &t)

	for k := 0; k < t; k++ {
		var n int
		fmt.Fscan(in, &n)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &a[i])
		}
		ans := make([]int, 0)
		for i := 1; i > 0; i++ {
			flag := true
			for j := 1; j < n; j++ {
				if a[j] > a[j+1] {
					flag = false
					break
				}
			}
			if flag {
				break
			}
			j := i & 1
			for j+1 <= n && a[j] < a[j+1] {
				j += 2
			}
			if j+1 <= n {
				ans = append(ans, j)
				a[j], a[j+1] = a[j+1], a[j]
			} else {
				ans = append(ans, j-2)
				a[j-2], a[j-1] = a[j-1], a[j-2]
			}
		}
		fmt.Fprintln(out, len(ans))
		for _, x := range ans {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}
}
