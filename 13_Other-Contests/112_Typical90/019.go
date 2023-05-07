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

	var a [444]int
	for i := 0; i < n+n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var d [444][222]int
	for i := 1; i <= n; i++ {
		for j := 0; i+i+j <= n+n; j++ {
			d[j][i] = d[j+1][i-1] + abs(a[j]-a[i+i+j-1])
			for k := 1; k < i; k++ {
				d[j][i] = min(d[j][i], d[j][k]+d[j+k+k][i-k])
			}
		}
	}
	fmt.Println(d[0][n])
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
