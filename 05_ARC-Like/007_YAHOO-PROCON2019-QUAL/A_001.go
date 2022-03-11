package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	if n >= k*2-1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
