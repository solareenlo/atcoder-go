package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	mini := 1000
	for i := 0; i < n-2; i++ {
		sub, _ := strconv.Atoi(s[i : i+3])
		mini = min(mini, abs(sub-753))
	}
	fmt.Println(mini)
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
