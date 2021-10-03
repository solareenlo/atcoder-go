package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	c, s, f := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&c[i], &s[i], &f[i])
	}

	for i := 0; i < n-1; i++ {
		res := s[i] + c[i]
		for j := i + 1; j < n-1; j++ {
			if res < s[j] {
				res = s[j]
			} else if res%f[j] != 0 {
				res += f[j] - (res % f[j])
			}
			res += c[j]
		}
		fmt.Println(res)
	}
	fmt.Println(0)
}
