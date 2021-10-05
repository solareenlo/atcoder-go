package main

import "fmt"

var (
	P [10]int = [10]int{0, 0, 0, 0, 0, 0, 1, 1, 1, 2}
	Q [10]int = [10]int{1, 1, 1, 2, 2, 3, 2, 2, 3, 3}
	R [10]int = [10]int{2, 3, 4, 3, 4, 4, 3, 4, 4, 4}
)

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	march := make([]int, 5)
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		switch s[0] {
		case 'M':
			march[0]++
		case 'A':
			march[1]++
		case 'R':
			march[2]++
		case 'C':
			march[3]++
		case 'H':
			march[4]++
		}
	}

	res := 0
	for i := 0; i < 10; i++ {
		res += march[P[i]] * march[Q[i]] * march[R[i]]
	}
	fmt.Println(res)
}
