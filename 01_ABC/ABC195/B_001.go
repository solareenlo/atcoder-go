package main

import "fmt"

func main() {
	var a, b, w int
	fmt.Scan(&a, &b, &w)

	w *= 1000
	mini := int(1e9)
	maxi := 0
	for i := 1; i < w+1; i++ {
		if a*i <= w && w <= b*i {
			mini = min(mini, i)
			maxi = max(maxi, i)
		}
	}

	if maxi != 0 {
		fmt.Println(mini, maxi)
	} else {
		fmt.Println("UNSATISFIABLE")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
