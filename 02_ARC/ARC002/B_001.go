package main

import "fmt"

func calLeapYear(y int) bool {
	ok := false
	if y%4 == 0 {
		ok = true
	}
	if y%100 == 0 {
		ok = false
	}
	if y%400 == 0 {
		ok = true
	}
	return ok
}

func main() {
	var y, m, d int
	fmt.Scanf("%d/%d/%d", &y, &m, &d)

	dMax := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for i := y; i <= 3000; i++ {
		mMin := 1
		if i == y {
			mMin = m
		}
		for j := mMin; j <= 12; j++ {
			dMin := 1
			if j == m {
				dMin = d
			}
			if j == 2 {
				leap := calLeapYear(y)
				if leap {
					dMax[2] = 29
				}
			}
			for k := dMin; k <= dMax[j]; k++ {
				if (float64(i)/float64(j))/float64(k) == float64((i/j)/k) {
					fmt.Printf("%02d/%02d/%02d\n", i, j, k)
					return
				}
			}
		}
	}
}
