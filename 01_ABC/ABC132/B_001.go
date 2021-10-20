package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}

	cnt := 0
	for i := 0; i < n-2; i++ {
		if (p[i] < p[i+1] && p[i+1] < p[i+2]) || (p[i] > p[i+1] && p[i+1] > p[i+2]) {
			cnt++
		}
	}
	fmt.Println(cnt)
}
