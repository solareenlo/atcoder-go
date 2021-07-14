package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 3)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Sort(sort.IntSlice(a))
	if a[0]-a[1] == a[1]-a[2] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
