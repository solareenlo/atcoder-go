package main

import (
	"fmt"
	"sort"
)

func main() {
	h := make([]int, 3)
	fmt.Scan(&h[0], &h[1], &h[2])
	sort.Ints(h)

	if h[0] == 5 && h[1] == 5 && h[2] == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
