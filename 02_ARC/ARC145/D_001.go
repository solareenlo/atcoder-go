package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, 100005)
	s := 0
	for i := 0; i <= n-2; i++ {
		for j, b := 0, 1; j < 15; j, b = j+1, b*3 {
			if i&(1<<j) != 0 {
				a[i+1] += b
			}
		}
		s += a[i+1]
	}

	r := (m - s) % n
	a[n] = (3*a[n-1]+n)/n*n + r
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d ", a[i]+(m-s-a[n])/n)
	}
}
