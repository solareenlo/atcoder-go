package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	if n == int64(int32(n)) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
