package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s, t, m int
	fmt.Fscan(in, &s, &t, &m)

	var x, y int
	a := make([][]int, 300005)
	for m > 0 {
		fmt.Fscan(in, &x, &y)
		a[x] = append(a[x], y-s)
		m--
	}

	var b [3005][3005]int
	for i := 1; i <= s; i++ {
		for _, x := range a[i] {
			for _, y := range a[i] {
				if x == y {
					continue
				}
				if b[x][y] != 0 {
					fmt.Printf("%d %d %d %d\n", x+s, y+s, b[x][y], i)
					os.Exit(0)
				}
				b[x][y] = i
			}
		}
	}
	fmt.Println(-1)
}
