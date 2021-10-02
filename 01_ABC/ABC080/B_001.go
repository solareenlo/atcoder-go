package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := strconv.Itoa(n)
	sum := 0
	for i := range s {
		sum += int(s[i] - '0')
	}
	if n%sum != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
