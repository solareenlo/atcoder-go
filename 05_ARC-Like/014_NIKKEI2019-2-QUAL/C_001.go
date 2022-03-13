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

	type pair struct{ x, y int }
	a := make([]pair, n)
	b := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i].x)
		a[i].y = i
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i].x)
		b[i].y = i
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].x < a[j].x
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i].x < b[j].x
	})

	p := make([]int, n)
	d := make([]int, n)
	o := true
	c := false
	for i := 0; i < n; i++ {
		if a[i].x > b[i].x {
			o = false
		}
		if i != 0 && a[i].x <= b[i-1].x {
			c = true
		}
		p[a[i].y] = b[i].y
	}

	z := 0
	for i := 0; i < n; i++ {
		if d[z] != 0 {
			c = true
		}
		d[z]++
		z = p[z]
	}

	if o && c {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
