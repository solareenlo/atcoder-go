package main

import "fmt"

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c, &k)

	cnt := 0
	for a >= b {
		b *= 2
		cnt++
	}
	for b >= c {
		c *= 2
		cnt++
	}

	if cnt <= k {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
