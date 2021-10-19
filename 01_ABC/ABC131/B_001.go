package main

import "fmt"

func main() {
	var n, l int
	fmt.Scan(&n, &l)

	sum := 0
	for i := 0; i < n; i++ {
		sum += l + i
	}

	mini := 1 << 60
	for i := 0; i < n; i++ {
		mini = min(mini, abs(l))
		l++
	}

	if l <= 0 {
		fmt.Println(sum + mini)
	} else {
		fmt.Println(sum - mini)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
