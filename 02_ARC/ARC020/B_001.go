package main

import "fmt"

func main() {
	var n, c int
	fmt.Scan(&n, &c)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
		a[i]--
	}

	mini := 1 << 60
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i != j {
				cnt := 0
				for k := 0; k < n; k += 2 {
					if i != a[k] {
						cnt++
					}
					if k+1 < n {
						if j != a[k+1] {
							cnt++
						}
					}
				}
				mini = min(mini, cnt)
			}
		}
	}

	fmt.Println(mini * c)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
