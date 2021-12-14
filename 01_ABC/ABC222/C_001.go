package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]string, n+n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	d := make([][2]int, n+n)
	for i := 0; i < n+n; i++ {
		d[i] = [2]int{0, i}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x := a[d[j+j][1]][i]
			y := a[d[j+j+1][1]][i]
			if x == 'G' && y == 'C' || x == 'C' && y == 'P' || x == 'P' && y == 'G' {
				d[j+j][0]--
			} else if x != y {
				d[j+j+1][0]--
			}
		}
		sort.Slice(d, func(i, j int) bool {
			return d[i][0] < d[j][0] || d[i][0] == d[j][0] && d[i][1] < d[j][1]
		})
	}

	for i := range d {
		fmt.Println(d[i][1] + 1)
	}
}
