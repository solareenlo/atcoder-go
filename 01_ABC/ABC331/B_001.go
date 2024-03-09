package main

import "fmt"

func main() {
	var n, s, m, l int
	fmt.Scan(&n, &s, &m, &l)
	r := 2147483647
	for i := 0; i*6 < 115; i++ {
		for j := 0; i*6+j*8 < 115; j++ {
			for k := 0; i*6+j*8+k*12 < 115; k++ {
				if i*6+j*8+k*12 >= n {
					r = min(r, i*s+j*m+k*l)
				}
			}
		}
	}
	fmt.Println(r)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
