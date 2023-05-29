package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	m := n - k
	j := 1
	for i := 0; i < k; i++ {
		fmt.Printf("%d ", j)
		j++
		if i%2 == 1 && m > 0 {
			j++
			m--
		}
	}
	fmt.Println()
}
