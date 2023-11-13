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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	b := make([]int, 0)
	t := 0
	c := make([]int, 0)
	for i := 0; i < n; i++ {
		if a[i] != t {
			b = append(b, a[i])
			c = append(c, 1)
			t = a[i]
		} else {
			c[len(c)-1]++
		}
	}

	r := n
	k := len(b)
	for i := 2; i < k; i++ {
		for b[i]-1 == b[i-1] && b[i]-2 == b[i-2] && c[i] > 0 && c[i-1] > 0 && c[i-2] > 0 {
			c[i]--
			c[i-1]--
			c[i-2]--
			r -= 3
		}
	}
	fmt.Println(r)
}
