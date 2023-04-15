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
	a := make([]int, 0)
	for i := 0; i < n; i++ {
		var b int
		fmt.Fscan(in, &b)
		a = append(a, b)
	}
	h, l := n, 0
	for h-l > 1 {
		m := (h + l) / 2
		t := 0
		sum := 1
		for i := 0; i < n-1; i++ {
			if a[i+1] == a[i] {
				sum++
			} else {
				sum = 1
			}
			if sum > m {
				t++
				sum = 0
			}
		}
		if t <= k {
			h = m
		} else {
			l = m
		}
	}
	fmt.Println(h)
}
