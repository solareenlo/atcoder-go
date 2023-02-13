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
	var a [1001][1001]int
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if a[i+1][j] >= a[i+1][j+1] {
				a[i][j] = a[i][j] + a[i+1][j]
			} else {
				a[i][j] = a[i][j] + a[i+1][j+1]
			}
		}
	}
	fmt.Println(a[0][0])
}
