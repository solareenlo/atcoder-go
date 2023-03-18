package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	var X [10]int
	var Y [10]int
	for i := 0; i < K; i++ {
		fmt.Scan(&X[i], &Y[i])
	}

	var x, y []int
	for i := 0; i < K; i++ {
		for d := -1; d <= 1; d++ {
			if X[i]+d > 0 && X[i]+d <= N {
				x = append(x, X[i]+d)
			}
			if Y[i]+d > 0 && Y[i]+d <= N {
				y = append(y, Y[i]+d)
			}
		}
	}

	ans := 0
	for _, x1 := range x {
		for _, x2 := range x {
			if x1 <= x2 {
				for _, y1 := range y {
					for _, y2 := range y {
						if y1 <= y2 {
							now := 0
							if (y2-y1)%2 == 0 && (x2-x1)%2 == 0 {
								if (x1+y1)%2 == 0 {
									now++
								} else {
									now--
								}
							}
							for i := 0; i < K; i++ {
								if x1 <= X[i] && X[i] <= x2 && y1 <= Y[i] && Y[i] <= y2 {
									if (X[i]+Y[i])%2 == 0 {
										now -= 2
									} else {
										now += 2
									}
								}
							}
							ans = max(ans, abs(now))
						}
					}
				}
			}
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
