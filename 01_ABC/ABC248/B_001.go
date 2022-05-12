package main

import "fmt"

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	cnt := 0
	for a < b {
		a *= k
		cnt++
	}
	fmt.Println(cnt)
}
