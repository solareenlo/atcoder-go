package main

import (
	"fmt"
)

func main() {
	var n, p int
	fmt.Scan(&n, &p)

	cnt := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if a < p {
			cnt++
		}
	}
	fmt.Println(cnt)
}
