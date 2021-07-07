package main

import "fmt"

func main() {
	var a, b, c, d int64
	fmt.Scan(&a, &b, &c, &d)
	diff := c*d - b
	if diff > 0 {
		fmt.Println((a + diff - 1) / diff)
	} else {
		fmt.Println(-1)
	}
}
