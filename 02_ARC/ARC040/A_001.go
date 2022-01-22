package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]string, n)
	for i := range s {
		fmt.Scan(&s[i])
	}

	cntR := 0
	cntB := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == 'R' {
				cntR++
			}
			if s[i][j] == 'B' {
				cntB++
			}
		}
	}

	if cntR == cntB {
		fmt.Println("DRAW")
	} else if cntR > cntB {
		fmt.Println("TAKAHASHI")
	} else {
		fmt.Println("AOKI")
	}
}
