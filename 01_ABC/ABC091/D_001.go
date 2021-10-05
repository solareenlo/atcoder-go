package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	bMod := make([]int, n)
	res := 0
	for j := 0; j < 32; j++ {
		mod := 1 << j
		for i := 0; i < n; i++ {
			bMod[i] = b[i] % (mod * 2)
		}
		sort.Ints(bMod)
		sum := 0
		for i := 0; i < n; i++ {
			aMod := a[i] % (mod * 2)
			b1 := binarySearch(bMod, 1*mod-aMod)
			b2 := binarySearch(bMod, 2*mod-aMod)
			b3 := binarySearch(bMod, 3*mod-aMod)
			sum += b2 - b1 + n - b3
		}
		if sum%2 == 1 {
			res += mod
		}
	}
	fmt.Println(res)
}

func binarySearch(a []int, x int) int {
	l, r := -1, len(a)
	for r-l > 1 {
		c := (l + r) / 2
		if a[c] >= x {
			r = c
		} else {
			l = c
		}
	}
	return r
}
