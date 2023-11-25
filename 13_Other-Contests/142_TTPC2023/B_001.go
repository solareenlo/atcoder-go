package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	P := make([]int, 13)
	P[0] = 1
	for i := 0; i < 12; i++ {
		P[i+1] = P[i] * 3
	}
	MX := P[12]
	s := -1
	g := -1
	mp := make([]bool, MX)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		mp[x] = true
		if i == 0 {
			s = x
		}
		if i == n-1 {
			g = x
		}
	}
	q := make([]int, 0)
	q = append(q, s)
	used := make([]bool, MX)
	var get func(int, int) int
	get = func(x, i int) int {
		x /= P[i]
		return x % 3
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		if x == g {
			fmt.Println("Yes")
			return
		}
		if used[x] {
			continue
		}
		used[x] = true
		for i := 0; i < 12; i++ {
			if get(x, i) != 2 {
				y := x + P[i]
				if !used[y] {
					q = append(q, y)
				}
			}
			if mp[x] {
				y := x - get(x, i)*P[i]
				if !used[y] {
					q = append(q, y)
				}
			}
		}
	}
	fmt.Println("No")
}
