package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i < 3501; i++ {
		for j := 1; j < 3501; j++ {
			d := 4*i*j - n*i - n*j
			e := n * i * j
			if 0 < d && e%d == 0 {
				fmt.Println(i, j, e/d)
				return
			}
		}
	}
}
