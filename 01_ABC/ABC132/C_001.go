package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	d := make([]int, n)
	for i := range d {
		fmt.Scan(&d[i])
	}
	sort.Ints(d)

	midL := d[len(d)/2-1]
	midR := d[len(d)/2]

	fmt.Println(midR - midL)
}
