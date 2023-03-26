package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	p := make([]Point, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&p[i].first, &p[i].second)
	}
	fmt.Println(3*(N-1) - len(convex_hull(p)))
}

type Point struct {
	first  int
	second int
}

func minus(a, b Point) Point {
	return Point{a.first - b.first, a.second - b.second}
}

func cross(a, b Point) int {
	return a.first*b.second - a.second*b.first
}

func convex_hull(p []Point) []Point {
	n := len(p)
	k := 0
	if n <= 2 {
		return p
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].first == p[j].first {
			return p[i].second < p[j].second
		}
		return p[i].first < p[j].first
	})
	ch := make([]Point, 2*n)
	for i := 0; i < n; i, k = i+1, k+1 {
		for k >= 2 && cross(minus(ch[k-1], ch[k-2]), minus(p[i], ch[k-1])) < 0 {
			k--
		}
		ch[k] = p[i]
	}
	for i, t := n-2, k+1; i >= 0; k, i = k+1, i-1 {
		for k >= t && cross(minus(ch[k-1], ch[k-2]), minus(p[i], ch[k-1])) < 0 {
			k--
		}
		ch[k] = p[i]
	}
	resize(&ch, k-1)
	return ch
}

func resize(a *[]Point, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, Point{0, 0})
		}
	}
}
