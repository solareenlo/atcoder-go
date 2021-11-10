package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)

	sum := 0
	for i := 0; i < n; i++ {
		sum += int(s[i] - '0')
	}

	if sum%9 != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
