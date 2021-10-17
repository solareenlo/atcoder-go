package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	v := make([]int, n)
	for i := range v {
		fmt.Scan(&v[i])
	}
	c := make([]int, n)
	for i := range c {
		fmt.Scan(&c[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		if v[i]-c[i] > 0 {
			res += v[i] - c[i]
		}
	}
	fmt.Println(res)
}
