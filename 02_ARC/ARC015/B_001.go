package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cnt := make([]int, 6)
	for i := 0; i < n; i++ {
		var M, m float64
		fmt.Scan(&M, &m)
		if M >= 35 {
			cnt[0]++
		}
		if M < 35 && M >= 30 {
			cnt[1]++
		}
		if M < 30 && M >= 25 {
			cnt[2]++
		}
		if m >= 25 {
			cnt[3]++
		}
		if m < 0 && M >= 0 {
			cnt[4]++
		}
		if M < 0 {
			cnt[5]++
		}
	}

	for i := 0; i < 6; i++ {
		fmt.Print(cnt[i])
		if i != 5 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
