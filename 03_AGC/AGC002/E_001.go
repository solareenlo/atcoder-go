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

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a[1:])))

	p := 1
	for p < n && a[p+1] >= p+1 {
		p++
	}

	q := p
	for q < n && a[q+1] == p {
		q++
	}

	if ((q-p)&1) != 0 || ((a[p]-p)&1) != 0 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
