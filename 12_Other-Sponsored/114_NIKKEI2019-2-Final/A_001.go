package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a [5005]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	ans := 0
	for m := 2; m < n; m++ {
		cntl, cntr := 0, 0
		for l := 1; l < m; l++ {
			if a[l] < a[m] {
				cntl++
			}
		}
		for r := m + 1; r <= n; r++ {
			if a[m] > a[r] {
				cntr++
			}
		}
		ans += cntl * cntr
	}
	fmt.Println(ans)
}
