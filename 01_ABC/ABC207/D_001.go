package main

import (
	"fmt"
	"reflect"
	"sort"
)

func cross(a, b, c [2]int) int {
	return (a[0]-c[0])*(b[1]-c[1]) - (a[1]-c[1])*(b[0]-c[0])
}

func dist(a, b [2]int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}

func main() {
	var n int
	fmt.Scan(&n)

	s := make([][2]int, n)
	t := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i][0], &s[i][1])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i][0], &t[i][1])
	}

	if n == 1 {
		fmt.Println("Yes")
		return
	} else if n == 2 {
		if dist(s[0], s[1]) == dist(t[0], t[1]) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
		return
	}

	ds := make([][3]int, n)
	for i := 0; i < n; i++ {
		ds[i][0] = dist(s[0], s[i])
		ds[i][1] = dist(s[1], s[i])
		ds[i][2] = cross(s[0], s[1], s[i])
	}
	sort.Slice(ds, func(i, j int) bool {
		if ds[i][0] == ds[j][0] {
			if ds[i][1] == ds[j][1] {
				return ds[i][2] < ds[j][2]
			}
			return ds[i][1] < ds[j][1]
		}
		return ds[i][0] < ds[j][0]
	})

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			dt := make([][3]int, n)
			for k := 0; k < n; k++ {
				dt[k][0] = dist(t[i], t[k])
				dt[k][1] = dist(t[j], t[k])
				dt[k][2] = cross(t[i], t[j], t[k])
			}
			sort.Slice(dt, func(i, j int) bool {
				if dt[i][0] == dt[j][0] {
					if dt[i][1] == dt[j][1] {
						return dt[i][2] < dt[j][2]
					}
					return dt[i][1] < dt[j][1]
				}
				return dt[i][0] < dt[j][0]
			})
			if reflect.DeepEqual(dt, ds) {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
