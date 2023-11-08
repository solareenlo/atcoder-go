package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type rou struct {
		n, k int
		a    [100]int
	}

	var a [38]rou
	for i := range a {
		a[i].k = 40
	}

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var c int
		fmt.Fscan(in, &c)
		for j := 0; j < c; j++ {
			var b int
			fmt.Fscan(in, &b)
			if c == a[b].k {
				a[b].a[a[b].n] = i
				a[b].n++
			} else if c < a[b].k {
				a[b].n = 1
				a[b].a[0] = i
				a[b].k = c
			}
		}
	}
	fmt.Fscan(in, &n)
	fmt.Println(a[n].n)
	for i := 0; i < a[n].n; i++ {
		fmt.Printf("%d ", a[n].a[i]+1)
	}
}
