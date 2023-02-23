package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var MonthDay = [20]int{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	var n int
	fmt.Fscan(in, &n)
	sum := make([]int, 15)
	for i := 1; i <= 12; i++ {
		sum[i] = sum[i-1] + MonthDay[i]
	}
	u := make([]bool, 370)
	for i := 1; i <= 366; i++ {
		if i%7 == 0 || i%7 == 1 {
			u[i] = true
		}
	}
	for i := 1; i <= n; i++ {
		var m, d int
		fmt.Fscanf(in, "\n%d/%d", &m, &d)
		if u[sum[m-1]+d] {
			for j := sum[m-1] + d + 1; j <= 366; j++ {
				if !u[j] {
					u[j] = true
					break
				}
			}
		} else {
			u[sum[m-1]+d] = true
		}
	}
	now := 0
	ans := 0
	for i := 1; i <= 366; i++ {
		if u[i] {
			now++
		} else {
			ans = max(ans, now)
			now = 0
		}
	}
	fmt.Println(max(ans, now))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
