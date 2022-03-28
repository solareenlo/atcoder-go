package main

import (
	"bufio"
	"fmt"
	"os"
)

func p(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	na := 0
	a := make([]int, 1010)
	for i := 2; i <= 8000; i++ {
		if p(i) {
			na++
			a[na] = i
		}
	}

	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if (i+j)%2 == 0 {
				fmt.Fprint(out, a[(i+j)/2]*a[(i-j+n)/2+n+1], " ")
			} else {
				fmt.Fprint(out, a[(i+j)/2]*a[(i-j+n+1)/2+n+1]*a[(i+j+1)/2]*a[(i-j+n-1)/2+n+1]+1, " ")
			}
		}
		fmt.Fprintln(out)
	}
}
