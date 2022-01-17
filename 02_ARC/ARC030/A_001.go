package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	if n/2 >= k {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
