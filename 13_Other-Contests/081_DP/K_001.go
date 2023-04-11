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
	var a [1 << 18]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	var f [1 << 18]bool
	for i := 1; i <= k; i++ {
		for j := 1; j <= n; j++ {
			if i >= a[j] && (!f[i-a[j]]) {
				f[i] = true
			}
		}
	}
	if f[k] {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
