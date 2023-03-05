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
	a := make([]int, 3000000)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	tmp := a[:n]
	sort.Ints(tmp)
	reverseOrderInt(tmp)

	r := n - 1
	for l := 0; l < r; l++ {
		a[r+1] = a[l] % a[r]
		if a[r+1] != 0 {
			r++
		}
	}
	fmt.Println(r)
}

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
