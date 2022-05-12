package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	x := [310]int{}
	y := [310]int{}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	if k == 1 {
		fmt.Println("Infinity")
		return
	}

	t := [310]int{}
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			s := 2
			for k := 1; k <= n; k++ {
				if i != k && j != k {
					if x[j] == x[k] && x[i] == x[k] || x[j] != x[k] && x[i] != x[j] && (y[k]-y[j])*(x[j]-x[i]) == (y[j]-y[i])*(x[k]-x[j]) {
						s++
					}
				}
			}
			t[s]++
		}
	}

	ans := 0
	for i := k; i <= n; i++ {
		ans += t[i] / (i * (i - 1) / 2)
	}
	fmt.Println(ans)
}
