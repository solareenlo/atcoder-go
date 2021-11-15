package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	mini := 100
	sum := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			var a int
			fmt.Scan(&a)
			sum += a
			mini = min(mini, a)
		}
	}

	fmt.Println(sum - mini*h*w)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
