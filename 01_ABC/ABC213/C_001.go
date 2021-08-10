package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var H, W, n int
	fmt.Scan(&H, &W, &n)

	a := make([]Card, n)
	b := make([]Card, n)
	var A, B int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A, &B)
		a[i].number = A
		a[i].index = i
		b[i].number = B
		b[i].index = i
	}
	sort.Sort(Cards(a))
	sort.Sort(Cards(b))

	c := make([]int, n)
	d := make([]int, n)
	ci, di := 1, 1
	for i := 0; i < n; i++ {
		c[a[i].index] = ci
		d[b[i].index] = di
		if i != n-1 {
			if a[i+1].number != a[i].number {
				ci++
			}
			if b[i+1].number != b[i].number {
				di++
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, c[i], d[i])
	}
}

type Card struct {
	number int
	index  int
}

type Cards []Card

func (a Cards) Len() int {
	return len(a)
}

func (a Cards) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Cards) Less(i, j int) bool {
	return a[i].number < a[j].number
}
