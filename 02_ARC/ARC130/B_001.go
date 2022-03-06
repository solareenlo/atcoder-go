package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w, c, q int
	fmt.Fscan(in, &h, &w, &c, &q)

	t := make([]int, q+1)
	a := make([]int, q+1)
	b := make([]int, q+1)
	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &t[i], &a[i], &b[i])
	}

	mp := [3]map[int]int{}
	for i := range mp {
		mp[i] = map[int]int{}
	}
	H, W := 0, 0
	ans := make([]int, 300300)
	for i := q; i > 0; i-- {
		T := t[i]
		A := a[i]
		B := b[i]
		if mp[T][A] != 0 {
			continue
		}
		mp[T][A] = 1
		if T == 1 {
			ans[B] += w - W
			H++
		} else {
			ans[B] += h - H
			W++
		}
	}
	for i := 1; i <= c; i++ {
		fmt.Fprint(out, ans[i], " ")
	}
}
