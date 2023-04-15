package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	isPossible := true
	for i := 0; i < n-1; i++ {
		if (s[i] == 'G' && s[i+1] == 'G') || (s[i] == 'R' && s[i+1] == 'B') {
			isPossible = false
		}
	}
	if s[0] != 'R' || s[n-1] != 'B' {
		isPossible = false
	}
	if isPossible {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
