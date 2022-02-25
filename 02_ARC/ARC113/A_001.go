package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)

	res := 0
	for i := 1; i <= k; i++ {
		for j := 1; j <= k/i; j++ {
			for l := 1; l <= k/i/j; l++ {
				res++
			}
		}
	}
	fmt.Println(res)
}
