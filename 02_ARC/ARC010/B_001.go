package main

import "fmt"

func main() {
	maxDay := map[int]int{}
	for i := 0; i < 12; i++ {
		if i == 1 {
			maxDay[i] = 29
		}
		if i == 3 || i == 5 || i == 8 || i == 10 {
			maxDay[i] = 30
		}
		if i == 0 || i == 2 || i == 4 || i == 6 || i == 7 || i == 9 || i == 11 {
			maxDay[i] = 31
		}
	}

	days := map[int]bool{}
	for i := 0; i < 366; i++ {
		days[i] = false
	}
	for i := 0; i < 366; i++ {
		if i%7 == 0 || i%7 == 6 {
			days[i] = true
		}
	}

	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var m, d int
		fmt.Scanf("%d/%d", &m, &d)
		m--
		d--
		sum := 0
		for j := 0; j < m; j++ {
			sum += maxDay[j]
		}
		sum += d
		for days[sum] && sum < 366 {
			sum++
		}
		if sum < 366 {
			days[sum] = true
		}
	}

	sum := 0
	ret := 0
	for i := 0; i < 366; i++ {
		if days[i] {
			sum++
			ret = max(ret, sum)
		} else {
			sum = 0
		}
	}

	fmt.Println(ret)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
