package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	for i := range s {
		fmt.Scan(&s[i])
	}

	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			fmt.Print(string(s[j][i]))
		}
		fmt.Println()
	}
}
