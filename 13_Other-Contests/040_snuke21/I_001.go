package main

import (
	"fmt"
	"sort"
)

func main() {

	const N = 1000000

	type pair struct {
		x, y int
	}

	a := make([]pair, 0)
	sum := 0
	for s := 2; sum+s <= N; s++ {
		for x := 1; x < s && sum+s <= N; x++ {
			if gcd(x, s) == 1 {
				a = append(a, pair{x, s - x})
				sum += s
			}
		}
	}

	a = append(a, pair{1, 0})
	a = append(a, pair{96, 97})

	sumx := 0
	for _, tmp := range a {
		sumx += tmp.x
	}

	sort.Slice(a, func(p, q int) bool {
		return a[p].x*a[q].y-a[q].x*a[p].y > 0
	})

	p := make([]pair, 0)
	x := N - sumx
	y := 0
	for t := 0; t < 4; t++ {
		for i := range a {
			if abs(a[i].x)+abs(a[i].y) == 193 && t%2 == 1 {
				a[i].x, a[i].y = -a[i].y, a[i].x
				continue
			}
			p = append(p, pair{x, y})
			x += a[i].x
			y += a[i].y
			a[i].x, a[i].y = -a[i].y, a[i].x
		}
	}

	var n int
	fmt.Scan(&n)

	if n > len(p) {
		fmt.Println("NO")
		return
	}

	fmt.Println("YES")
	for i := 0; i < n; i++ {
		fmt.Println(p[i].x, p[i].y)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
