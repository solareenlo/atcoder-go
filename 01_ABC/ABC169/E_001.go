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

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}
	sort.Ints(a)
	sort.Ints(b)

	var l, r int
	if n%2 != 0 {
		l = a[n/2]
		r = b[n/2]
	} else {
		l = a[n/2-1] + a[n/2]
		r = b[n/2-1] + b[n/2]
	}

	fmt.Println(r - l + 1)
}
