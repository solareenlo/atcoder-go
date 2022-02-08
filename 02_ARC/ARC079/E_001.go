package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	f := true
	for f {
		f = false
		for i := 1; i <= n; i++ {
			if a[i] >= n {
				f = true
				for j := 0; j <= n; j++ {
					if j != i {
						a[j] += a[i] / n
					}
				}
				a[i] %= n
			}
		}
	}
	fmt.Println(a[0])
}
