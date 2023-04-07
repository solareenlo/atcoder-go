package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var S string
	var gx, gy int
	fmt.Fscan(in, &S, &gx, &gy)
	inv := []int{0, 1, 2, 3}
	dx := [4]int{1, -1, 0, 0}
	dy := [4]int{0, 0, 1, -1}
	for nextPermutation(sort.IntSlice(inv)) {
		a, b := 0, 0
		for _, p := range S {
			if gx == a && gy == b {
				fmt.Println("Yes")
				return
			}
			a += dy[inv[p-'W']]
			b += dx[inv[p-'W']]
			if gx == a && gy == b {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
