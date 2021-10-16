package main

import (
	"fmt"
	"sort"
)

func main() {
	var x, y, z, K int
	fmt.Scan(&x, &y, &z, &K)

	a := make([]int, x)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	b := make([]int, y)
	for i := range b {
		fmt.Scan(&b[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(b)))

	c := make([]int, z)
	for i := range c {
		fmt.Scan(&c[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(c)))

	abc := make([]int, 0)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			for k := 0; k < z; k++ {
				if (i+1)*(j+1)*(k+1) > K {
					break
				}
				abc = append(abc, a[i]+b[j]+c[k])
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(abc)))

	for i := 0; i < K; i++ {
		fmt.Println(abc[i])
	}
}
