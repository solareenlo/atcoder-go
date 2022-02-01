package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w, n int
	fmt.Fscan(in, &h, &w, &n)

	inf := int(1e9 + 7)
	k := 0
	zo := (h - 2) * (w - 2)
	a := make([]int, (n+1)*16)
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		for i := 0; i <= 2; i++ {
			for j := 0; j <= 2; j++ {
				if x-i >= 1 && x-i <= h-2 && y-j >= 1 && y-j <= w-2 {
					a[k] = inf*(x-i) + (y - j)
					k++
				}
			}
		}
	}
	sort.Ints(a[:k])

	sum := make([]int, 20)
	t := 1
	for i := 0; i < k; i++ {
		if a[i] == a[i+1] {
			t++
		} else {
			sum[t]++
			t = 1
			zo--
		}
	}

	fmt.Println(zo)
	for i := 1; i <= 9; i++ {
		fmt.Println(sum[i])
	}
}
