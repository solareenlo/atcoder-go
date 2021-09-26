package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	m := map[int]int{}
	cnt, t := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		if t >= 3200 {
			cnt++
		} else {
			m[t/400]++
		}
	}

	if len(m) == 0 {
		fmt.Print(1, " ", cnt, "\n")
	} else {
		fmt.Print(len(m), " ", len(m)+cnt, "\n")
	}
}
