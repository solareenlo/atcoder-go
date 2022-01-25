package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	t := false
	if n <= a {
		t = true
	} else if n > a {
		if a == b {
			if n%(a+1) != 0 {
				t = true
			}
		} else {
			if a > b {
				t = true
			}
		}
	}

	if t {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}
