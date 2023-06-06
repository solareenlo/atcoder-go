package main

import "fmt"

func main() {
	d10 := []int{0, 3, 6, 8, 8, 8, 8, 9}
	d11 := []int{10, 9, 8, 9}
	d2 := []int{18, 20, 18, 16}
	var M, N int
	fmt.Scan(&M, &N)
	if M == 1 || N == 1 {
		k := max(M, N)
		if k < 8 {
			fmt.Println(d10[k])
		} else {
			fmt.Println(d11[k%4])
		}
	} else {
		if M == 3 && N == 3 {
			fmt.Println(28)
		} else if M+N == 6 {
			fmt.Println(16)
		} else {
			fmt.Println(d2[(M+N)%4])
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
