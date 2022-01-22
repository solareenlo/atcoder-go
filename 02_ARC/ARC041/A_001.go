package main

import "fmt"

func main() {
	var x, y, k int
	fmt.Scan(&x, &y, &k)

	ans := 0
	if k <= y {
		ans += x + k
	} else {
		ans += y + (x + y) - k
	}

	fmt.Println(ans)
}
