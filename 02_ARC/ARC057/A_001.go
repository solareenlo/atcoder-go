package main

import "fmt"

func main() {
	var a, k int
	fmt.Scan(&a, &k)

	if k == 0 {
		fmt.Println(int(2e12) - a)
	} else {
		cnt := 0
		for a < int(2e12) {
			a += 1 + k*a
			cnt++
		}
		fmt.Println(cnt)
	}
}
