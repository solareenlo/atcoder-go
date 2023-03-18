package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	r := make([][2]int, 0)
	for i := n - 1; i >= 0; i-- {
		for a[i] != i+1 {
			r = append(r, [2]int{i + 1 - a[i], i + 1})
			a[i-a[i]], a[i] = a[i], a[i-a[i]]
		}
	}
	fmt.Println(len(r))
	for _, s := range r {
		fmt.Println(s[0], s[1])
	}
}
