package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	s += " "

	c := [4]string{"A", "B", "X", "Y"}

	res := n
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for a := 0; a < 4; a++ {
				for b := 0; b < 4; b++ {
					L := c[i] + c[j]
					R := c[a] + c[b]
					cnt := 0
					for k := 0; k < n; k++ {
						sub := s[k : k+2]
						if k != n-1 && (sub == L || sub == R) {
							k++
						}
						cnt++
					}
					res = min(res, cnt)
				}
			}
		}
	}

	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
