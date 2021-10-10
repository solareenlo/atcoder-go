package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cnt := 0
	for i := 1; i <= n; i += 2 {
		div := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				div++
			}
		}
		if div == 8 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
